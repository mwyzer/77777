package inbox

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

// ── Migrations ──

func (r *Repository) RunMigrations(ctx context.Context) error {
	migration := `
	CREATE EXTENSION IF NOT EXISTS "pgcrypto";

	CREATE TABLE IF NOT EXISTS customers (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		phone VARCHAR(50),
		email VARCHAR(255),
		provider VARCHAR(20) NOT NULL CHECK (provider IN ('whatsapp', 'telegram')),
		provider_id VARCHAR(255),
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		UNIQUE (provider, provider_id)
	);

	CREATE INDEX IF NOT EXISTS idx_customers_phone ON customers (phone);
	CREATE INDEX IF NOT EXISTS idx_customers_provider ON customers (provider);

	CREATE TABLE IF NOT EXISTS conversations (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
		channel VARCHAR(20) NOT NULL CHECK (channel IN ('whatsapp', 'telegram')),
		status VARCHAR(20) NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'closed')),
		last_message_at TIMESTAMPTZ,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_conversations_customer_id ON conversations (customer_id);
	CREATE INDEX IF NOT EXISTS idx_conversations_status ON conversations (status);
	CREATE INDEX IF NOT EXISTS idx_conversations_channel ON conversations (channel);
	CREATE INDEX IF NOT EXISTS idx_conversations_last_message ON conversations (last_message_at DESC);

	CREATE TABLE IF NOT EXISTS messages (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		conversation_id UUID NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
		sender_type VARCHAR(10) NOT NULL CHECK (sender_type IN ('customer', 'agent')),
		content TEXT NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'sent' CHECK (status IN ('sent', 'delivered', 'read', 'failed')),
		provider_message_id VARCHAR(255),
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_messages_conversation_id ON messages (conversation_id);
	CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages (created_at);
	CREATE INDEX IF NOT EXISTS idx_messages_provider_message_id ON messages (provider_message_id);
	`

	_, err := r.db.Exec(ctx, migration)
	if err != nil {
		return fmt.Errorf("failed to run core inbox migration: %w", err)
	}

	return nil
}

// ── Customers ──

func (r *Repository) CreateCustomer(ctx context.Context, customer *Customer) error {
	query := `
		INSERT INTO customers (name, phone, email, provider, provider_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(ctx, query,
		customer.Name, customer.Phone, customer.Email,
		customer.Provider, customer.ProviderID,
	).Scan(&customer.ID, &customer.CreatedAt, &customer.UpdatedAt)
}

func (r *Repository) FindCustomerByProvider(ctx context.Context, provider, providerID string) (*Customer, error) {
	query := `
		SELECT id, name, phone, email, provider, provider_id, created_at, updated_at
		FROM customers
		WHERE provider = $1 AND provider_id = $2
	`
	c := &Customer{}
	err := r.db.QueryRow(ctx, query, provider, providerID).Scan(
		&c.ID, &c.Name, &c.Phone, &c.Email,
		&c.Provider, &c.ProviderID, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find customer by provider: %w", err)
	}
	return c, nil
}

func (r *Repository) FindCustomerByID(ctx context.Context, id string) (*Customer, error) {
	query := `
		SELECT id, name, phone, email, provider, provider_id, created_at, updated_at
		FROM customers
		WHERE id = $1
	`
	c := &Customer{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Phone, &c.Email,
		&c.Provider, &c.ProviderID, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find customer by id: %w", err)
	}
	return c, nil
}

func (r *Repository) ListCustomers(ctx context.Context, params CustomerFilterParams) ([]Customer, int, error) {
	page, limit := defaultPagination(params.Page, params.Limit)
	off := offset(page, limit)

	// Build WHERE clause
	where := "WHERE 1=1"
	args := []interface{}{}
	argIdx := 1

	if params.Search != "" {
		where += fmt.Sprintf(" AND (name ILIKE $%d OR phone ILIKE $%d OR email ILIKE $%d)", argIdx, argIdx+1, argIdx+2)
		searchPattern := "%" + params.Search + "%"
		args = append(args, searchPattern, searchPattern, searchPattern)
		argIdx += 3
	}
	if params.Provider != "" {
		where += fmt.Sprintf(" AND provider = $%d", argIdx)
		args = append(args, params.Provider)
		argIdx++
	}

	// Count total
	countQuery := "SELECT COUNT(*) FROM customers " + where
	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count customers: %w", err)
	}

	// Fetch page
	selectQuery := fmt.Sprintf(`
		SELECT id, name, phone, email, provider, provider_id, created_at, updated_at
		FROM customers
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, where, argIdx, argIdx+1)
	args = append(args, limit, off)

	rows, err := r.db.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list customers: %w", err)
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Email,
			&c.Provider, &c.ProviderID, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan customer: %w", err)
		}
		customers = append(customers, c)
	}

	return customers, total, nil
}

// ── Conversations ──

func (r *Repository) CreateConversation(ctx context.Context, conv *Conversation) error {
	query := `
		INSERT INTO conversations (customer_id, channel, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRow(ctx, query,
		conv.CustomerID, conv.Channel, conv.Status,
	).Scan(&conv.ID, &conv.CreatedAt, &conv.UpdatedAt)
}

func (r *Repository) FindConversationByCustomerAndChannel(ctx context.Context, customerID, channel string) (*Conversation, error) {
	query := `
		SELECT id, customer_id, channel, status, last_message_at, created_at, updated_at
		FROM conversations
		WHERE customer_id = $1 AND channel = $2 AND status = 'open'
		ORDER BY created_at DESC
		LIMIT 1
	`
	c := &Conversation{}
	err := r.db.QueryRow(ctx, query, customerID, channel).Scan(
		&c.ID, &c.CustomerID, &c.Channel, &c.Status,
		&c.LastMessageAt, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find conversation: %w", err)
	}
	return c, nil
}

func (r *Repository) FindConversationByID(ctx context.Context, id string) (*Conversation, error) {
	query := `
		SELECT id, customer_id, channel, status, last_message_at, created_at, updated_at
		FROM conversations
		WHERE id = $1
	`
	c := &Conversation{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.CustomerID, &c.Channel, &c.Status,
		&c.LastMessageAt, &c.CreatedAt, &c.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find conversation by id: %w", err)
	}
	return c, nil
}

func (r *Repository) ListConversations(ctx context.Context, params ConversationFilterParams) ([]Conversation, int, error) {
	page, limit := defaultPagination(params.Page, params.Limit)
	off := offset(page, limit)

	where := "WHERE 1=1"
	args := []interface{}{}
	argIdx := 1

	if params.Status != "" {
		where += fmt.Sprintf(" AND c.status = $%d", argIdx)
		args = append(args, params.Status)
		argIdx++
	}
	if params.Channel != "" {
		where += fmt.Sprintf(" AND c.channel = $%d", argIdx)
		args = append(args, params.Channel)
		argIdx++
	}
	if params.Search != "" {
		where += fmt.Sprintf(" AND (cu.name ILIKE $%d OR cu.phone ILIKE $%d)", argIdx, argIdx+1)
		searchPattern := "%" + params.Search + "%"
		args = append(args, searchPattern, searchPattern)
		argIdx += 2
	}

	// Count
	countQuery := `
		SELECT COUNT(*)
		FROM conversations c
		JOIN customers cu ON cu.id = c.customer_id
	` + where
	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count conversations: %w", err)
	}

	// Fetch page
	selectQuery := fmt.Sprintf(`
		SELECT c.id, c.customer_id, c.channel, c.status, c.last_message_at, c.created_at, c.updated_at,
			   cu.id, cu.name, cu.phone, cu.email, cu.provider, cu.provider_id, cu.created_at, cu.updated_at
		FROM conversations c
		JOIN customers cu ON cu.id = c.customer_id
		%s
		ORDER BY COALESCE(c.last_message_at, c.created_at) DESC
		LIMIT $%d OFFSET $%d
	`, where, argIdx, argIdx+1)
	args = append(args, limit, off)

	rows, err := r.db.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list conversations: %w", err)
	}
	defer rows.Close()

	var conversations []Conversation
	for rows.Next() {
		var c Conversation
		var cu Customer
		if err := rows.Scan(
			&c.ID, &c.CustomerID, &c.Channel, &c.Status, &c.LastMessageAt, &c.CreatedAt, &c.UpdatedAt,
			&cu.ID, &cu.Name, &cu.Phone, &cu.Email, &cu.Provider, &cu.ProviderID, &cu.CreatedAt, &cu.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("scan conversation: %w", err)
		}
		c.Customer = &cu
		conversations = append(conversations, c)
	}

	return conversations, total, nil
}

func (r *Repository) ListConversationsByCustomerID(ctx context.Context, customerID string) ([]Conversation, error) {
	query := `
		SELECT id, customer_id, channel, status, last_message_at, created_at, updated_at
		FROM conversations
		WHERE customer_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		return nil, fmt.Errorf("list conversations by customer: %w", err)
	}
	defer rows.Close()

	var conversations []Conversation
	for rows.Next() {
		var c Conversation
		if err := rows.Scan(
			&c.ID, &c.CustomerID, &c.Channel, &c.Status,
			&c.LastMessageAt, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan conversation: %w", err)
		}
		conversations = append(conversations, c)
	}

	return conversations, nil
}

func (r *Repository) UpdateConversationLastMessage(ctx context.Context, conversationID string, at time.Time) error {
	query := `
		UPDATE conversations
		SET last_message_at = $2, updated_at = NOW()
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, conversationID, at)
	return err
}

// ── Messages ──

func (r *Repository) CreateMessage(ctx context.Context, msg *Message) error {
	query := `
		INSERT INTO messages (conversation_id, sender_type, content, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	return r.db.QueryRow(ctx, query,
		msg.ConversationID, msg.SenderType, msg.Content, msg.Status,
	).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *Repository) ListMessages(ctx context.Context, conversationID string, page, limit int) ([]Message, int, error) {
	page, limit = defaultPagination(page, limit)
	off := offset(page, limit)

	// Count
	var total int
	if err := r.db.QueryRow(ctx,
		"SELECT COUNT(*) FROM messages WHERE conversation_id = $1", conversationID,
	).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count messages: %w", err)
	}

	// Fetch page (oldest first for chat view)
	rows, err := r.db.Query(ctx, `
		SELECT id, conversation_id, sender_type, content, status, provider_message_id, created_at
		FROM messages
		WHERE conversation_id = $1
		ORDER BY created_at ASC
		LIMIT $2 OFFSET $3
	`, conversationID, limit, off)
	if err != nil {
		return nil, 0, fmt.Errorf("list messages: %w", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.ConversationID, &m.SenderType,
			&m.Content, &m.Status, &m.ProviderMessageID, &m.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan message: %w", err)
		}
		messages = append(messages, m)
	}

	return messages, total, nil
}
