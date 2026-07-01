# Phase 05 — Telegram Integration

## Objective

Integrasi Telegram Bot API: menerima webhook pesan masuk dari Telegram, menyimpan ke database, dan membalas pesan via Bot API.

## Acceptance Criteria

- [ ] `POST /api/webhooks/telegram` — menerima update dari Telegram
- [ ] Webhook payload divalidasi dan diparse
- [ ] Idempotency check via Redis (berdasarkan `update_id`)
- [ ] Customer otomatis dibuat/ditemukan berdasarkan Telegram user
- [ ] Conversation otomatis dibuat/ditemukan
- [ ] Message tersimpan ke `messages` table
- [ ] Telegram provider bisa mengirim pesan balasan via Bot API (`sendMessage`)
- [ ] Webhook endpoint public (tidak perlu JWT)

## Files to Create

- `phases/phase-05-telegram.md`
- `backend/internal/providers/telegram/telegram.go` — Telegram Bot API client
- `backend/internal/modules/webhook/handler.go` — Webhook HTTP handler
- `backend/internal/modules/webhook/service.go` — Webhook processing logic

## Files to Change

- `backend/internal/config/config.go` — Tambah `TelegramBotToken`
- `backend/cmd/api/main.go` — Register `/api/webhooks/telegram` route

## Test Commands

```bash
# Build
go build ./...

# Run
go run ./cmd/api

# Simulate Telegram webhook
curl -X POST http://localhost:8080/api/webhooks/telegram \
  -H "Content-Type: application/json" \
  -d '{
    "update_id": 123456789,
    "message": {
      "message_id": 1,
      "chat": {
        "id": 987654321,
        "first_name": "Test",
        "last_name": "User",
        "username": "testuser"
      },
      "text": "Hello from Telegram!",
      "date": 1719840000
    }
  }'
```
