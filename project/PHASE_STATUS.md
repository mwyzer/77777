# PHASE_STATUS.md

# Phase Status — Customer Communication Dashboard

| Phase | Name                  | Status      | Notes                                                                  |
| ----- | --------------------- | ----------- | ---------------------------------------------------------------------- |
| 01    | Project Setup         | PASSED      | Setup Golang, PostgreSQL, Redis, MinIO, Docker Compose — health endpoint OK |
| 02    | Authentication        | PASSED      | JWT login, bcrypt hash, protected route, default admin, /api/auth/*      |
| 03    | Core Inbox            | TODO        | Customer, conversation, message, inbox API                             |
| 04    | React Dashboard Base  | TODO        | Login UI, layout, protected route, inbox UI                            |
| 05    | Telegram Integration  | TODO        | Telegram webhook, receive message, reply message                       |
| 06    | Redis Queue Worker    | TODO        | Queue send message, retry, status update                               |
| 07    | MinIO Attachment      | TODO        | Presigned upload, signed download, attachment metadata                 |
| 08    | WhatsApp Integration  | TODO        | WhatsApp webhook, provider interface, send message                     |
| 09    | Template Message      | TODO        | CRUD template and template picker                                      |
| 10    | Auto-Reply Keyword    | TODO        | Keyword rule, auto-send reply through queue                            |
| 11    | Realtime Inbox        | TODO        | Redis Pub/Sub and SSE/WebSocket                                        |
| 12    | Dashboard Summary     | TODO        | Message statistics and Redis cache                                     |
| 13    | Kubernetes Deployment | TODO        | Manifest basic for backend, frontend, worker, Redis, MinIO, PostgreSQL |
| 14    | Final Review          | TODO        | Refactor, security review, performance check                           |

## Status Options

```text
TODO
IN_PROGRESS
PASSED
FAILED
SKIPPED
```

## Rule

Only one phase may be `IN_PROGRESS` at a time.

Do not start the next phase until the current phase is `PASSED`.
