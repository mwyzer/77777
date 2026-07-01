package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client wraps the Telegram Bot API.
type Client struct {
	botToken  string
	baseURL   string
	httpClient *http.Client
}

// NewClient creates a new Telegram Bot API client.
func NewClient(botToken string) *Client {
	return &Client{
		botToken: botToken,
		baseURL:  fmt.Sprintf("https://api.telegram.org/bot%s", botToken),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SendMessage sends a text message to a Telegram chat.
// Returns the Telegram message ID on success.
func (c *Client) SendMessage(chatID int64, text string) (int64, error) {
	if c.botToken == "" || c.botToken == "your-telegram-bot-token" {
		return 0, fmt.Errorf("TELEGRAM_BOT_TOKEN not configured")
	}

	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/sendMessage", c.baseURL)
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("telegram API call failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("telegram API returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		OK     bool `json:"ok"`
		Result struct {
			MessageID int64 `json:"message_id"`
		} `json:"result"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}

	if !result.OK {
		return 0, fmt.Errorf("telegram API returned ok=false")
	}

	return result.Result.MessageID, nil
}

// ── Webhook Payload Types ──

// Update represents a Telegram webhook update.
type Update struct {
	UpdateID int64   `json:"update_id"`
	Message  *Message `json:"message"`
}

// Message represents a Telegram message.
type Message struct {
	MessageID int64 `json:"message_id"`
	Chat      Chat  `json:"chat"`
	Text      string `json:"text"`
	Date      int64  `json:"date"`
}

// Chat represents a Telegram chat/user.
type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
