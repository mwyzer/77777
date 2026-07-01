# Phase 03 Report — Core Inbox

**Status:** PASSED

## Summary

Modul Core Inbox berhasil diimplementasi. Mencakup 3 tabel utama (`customers`, `conversations`, `messages`) dengan migration otomatis, 6 endpoint API yang dilindungi JWT, pagination, search, dan filter.

## Files Created

- `phases/phase-03-core-inbox.md` — Phase definition & acceptance criteria
- `backend/migrations/002_create_core_inbox.sql` — SQL migration file
- `backend/internal/modules/inbox/model.go` — Customer, Conversation, Message structs
- `backend/internal/modules/inbox/dto.go` — Request/Response DTOs + pagination helpers
- `backend/internal/modules/inbox/repository.go` — Database queries + RunMigrations
- `backend/internal/modules/inbox/service.go` — Business logic
- `backend/internal/modules/inbox/handler.go` — HTTP handlers (6 endpoints)

## Files Changed

- `backend/cmd/api/main.go` — Import inbox module, run migrations, register 6 protected routes
- `project/PHASE_STATUS.md` — Phase 03: TODO → IN_PROGRESS → PASSED

## Commands Executed

```bash
go build ./...    # PASS — no errors
go vet ./...      # PASS — no issues
```

## Test Result

| Test                    | Result                            |
| ----------------------- | --------------------------------- |
| `go build ./...`        | PASS                              |
| `go vet ./...`          | PASS                              |
| Docker integration test | SKIP (Docker Desktop not running) |

## API Endpoints Registered

| Method | Path                                    | Description                     |
| ------ | --------------------------------------- | ------------------------------- |
| GET    | `/api/inbox/conversations`              | List conversations (paginated)  |
| GET    | `/api/inbox/conversations/:id`          | Conversation detail + customer  |
| GET    | `/api/inbox/conversations/:id/messages` | Messages in conversation        |
| POST   | `/api/inbox/conversations/:id/messages` | Send reply (agent)              |
| GET    | `/api/inbox/customers`                  | List/search customers           |
| GET    | `/api/inbox/customers/:id`              | Customer detail + conversations |

## Database Tables Created

- `customers` — with indexes on phone, provider, unique(provider, provider_id)
- `conversations` — FK to customers, indexes on customer_id, status, channel, last_message_at
- `messages` — FK to conversations, indexes on conversation_id, created_at, provider_message_id

## Known Issues

- Docker Desktop tidak running saat test — integration test via `curl` tidak bisa dijalankan. Jalankan `docker compose up -d --build` secara manual lalu test endpoint di atas.

## Next Phase

Phase 04 — React Dashboard Base (Login UI, layout, protected route, inbox UI)
