# Phase 03 — Core Inbox

## Objective

Membuat modul Core Inbox: Customer, Conversation, Message, dan inbox API untuk membaca dan membalas pesan.

## Acceptance Criteria

- [ ] Table `customers`, `conversations`, `messages` berhasil dibuat via migration
- [ ] `GET /api/inbox/conversations` — list conversations (paginated, filter by status/channel)
- [ ] `GET /api/inbox/conversations/:id` — detail conversation + customer info
- [ ] `GET /api/inbox/conversations/:id/messages` — list messages in conversation (paginated)
- [ ] `POST /api/inbox/conversations/:id/messages` — send reply message (agent → customer)
- [ ] `GET /api/inbox/customers` — list/search customers (paginated)
- [ ] `GET /api/inbox/customers/:id` — customer detail + conversations
- [ ] Semua endpoint dilindungi JWT middleware
- [ ] Response menggunakan format JSON standar

## Database Schema

### customers

| Column      | Type         | Notes                         |
| ----------- | ------------ | ----------------------------- |
| id          | UUID PK      | gen_random_uuid()             |
| name        | VARCHAR(255) | NOT NULL                      |
| phone       | VARCHAR(50)  | nullable                      |
| email       | VARCHAR(255) | nullable                      |
| provider    | VARCHAR(20)  | whatsapp, telegram            |
| provider_id | VARCHAR(255) | nullable, unique per provider |
| created_at  | TIMESTAMPTZ  | NOT NULL DEFAULT NOW()        |
| updated_at  | TIMESTAMPTZ  | NOT NULL DEFAULT NOW()        |

### conversations

| Column          | Type        | Notes                  |
| --------------- | ----------- | ---------------------- |
| id              | UUID PK     | gen_random_uuid()      |
| customer_id     | UUID FK     | → customers(id)        |
| channel         | VARCHAR(20) | whatsapp, telegram     |
| status          | VARCHAR(20) | open, closed           |
| last_message_at | TIMESTAMPTZ | nullable               |
| created_at      | TIMESTAMPTZ | NOT NULL DEFAULT NOW() |
| updated_at      | TIMESTAMPTZ | NOT NULL DEFAULT NOW() |

### messages

| Column              | Type         | Notes                         |
| ------------------- | ------------ | ----------------------------- |
| id                  | UUID PK      | gen_random_uuid()             |
| conversation_id     | UUID FK      | → conversations(id)           |
| sender_type         | VARCHAR(10)  | customer, agent               |
| content             | TEXT         | NOT NULL                      |
| status              | VARCHAR(20)  | sent, delivered, read, failed |
| provider_message_id | VARCHAR(255) | nullable                      |
| created_at          | TIMESTAMPTZ  | NOT NULL DEFAULT NOW()        |

## API Endpoints

| Method | Path                                  | Auth | Description                         |
| ------ | ------------------------------------- | ---- | ----------------------------------- |
| GET    | /api/inbox/conversations              | Yes  | List conversations (paginated)      |
| GET    | /api/inbox/conversations/:id          | Yes  | Get conversation detail             |
| GET    | /api/inbox/conversations/:id/messages | Yes  | List messages in conversation       |
| POST   | /api/inbox/conversations/:id/messages | Yes  | Send reply message (agent)          |
| GET    | /api/inbox/customers                  | Yes  | List/search customers (paginated)   |
| GET    | /api/inbox/customers/:id              | Yes  | Get customer detail + conversations |

## Test Commands

```bash
# Build check
go build ./...

# Run server
go run ./cmd/api

# Test endpoints (after docker compose up -d and login)
TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"admin123"}' | jq -r '.data.token')

curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/inbox/conversations
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/inbox/customers
```
