# Customer Communication Dashboard

Unified dashboard untuk mengelola pesan customer dari **WhatsApp** dan **Telegram** dalam satu inbox.

## Tech Stack

| Layer          | Teknologi                          |
| -------------- | ---------------------------------- |
| Backend        | Go + Gin                           |
| Frontend       | React + Vite + TypeScript          |
| Database       | PostgreSQL 16                      |
| Cache / Queue  | Redis 7                            |
| Object Storage | MinIO                              |
| Container      | Docker + Docker Compose            |
| Orchestration  | Kubernetes (production)            |

## MVP Features

- [x] Project Setup — Go backend, Docker Compose, health check
- [x] Authentication — JWT login, bcrypt hash, protected routes
- [ ] Unified Inbox — Customer, conversation, message APIs
- [ ] Telegram Integration — Webhook receive & reply
- [ ] WhatsApp Integration — Webhook receive & reply
- [ ] Auto-Reply — Keyword-based automatic response
- [ ] Message Templates — CRUD + template picker
- [ ] Attachment — Presigned upload/download via MinIO
- [ ] Redis Queue — Outgoing message queue with retry
- [ ] Realtime Inbox — Redis Pub/Sub + SSE/WebSocket
- [ ] Dashboard Summary — Message statistics + Redis cache

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- [Go](https://go.dev/dl/) 1.21+ (for local dev)
- [Node.js](https://nodejs.org/) 18+ (for frontend dev)

## Quick Start

```bash
# 1. Clone repository
git clone https://github.com/mwyzer/77777.git
cd 77777

# 2. Copy environment file
cp .env.example .env

# 3. Start all services
docker compose up -d

# 4. Verify
curl http://localhost:8080/health
```

Services:

| Service          | URL                      |
| ---------------- | ------------------------ |
| Backend API      | http://localhost:8080    |
| MinIO Console    | http://localhost:9001    |
| PostgreSQL       | localhost:5432           |
| Redis            | localhost:6379           |

## Project Structure

```
├── backend/
│   ├── cmd/api/main.go           # API entry point
│   ├── internal/
│   │   ├── config/               # App configuration
│   │   ├── database/             # PostgreSQL, Redis, MinIO clients
│   │   ├── middleware/            # Auth middleware (JWT)
│   │   ├── modules/              # Modular features (auth, inbox, etc.)
│   │   └── response/             # Standard JSON response helpers
│   ├── migrations/               # SQL migration files
│   └── Dockerfile
├── docker-compose.yml            # Local dev stack
├── docs/                         # Architecture, API contract, BRD, SRS
├── phases/                       # Phase-by-phase implementation plans
├── project/                      # Project management (rules, status, reports)
├── scripts/                      # Check and test scripts
└── prompts/                      # AI coding agent prompt templates
```

## Environment Variables

Copy `.env.example` to `.env` and adjust:

```env
SERVER_PORT=8080
DATABASE_URL=postgres://dashboard:dashboard@postgres:5432/dashboard?sslmode=disable
REDIS_ADDR=redis:6379
REDIS_PASS=
MINIO_ENDPOINT=minio:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_BUCKET=chat-media
MINIO_USE_SSL=false
JWT_SECRET=change-me-in-production
```

## API Endpoints

| Method | Path                  | Auth     | Description          |
| ------ | --------------------- | -------- | -------------------- |
| GET    | `/health`             | No       | Health check         |
| POST   | `/api/auth/login`     | No       | User login           |
| POST   | `/api/auth/register`  | No       | User registration    |
| GET    | `/api/auth/me`        | Yes      | Current user profile |

## Development Workflow

Project dikerjakan secara bertahap mengikuti **phase** di `phases/`. Lihat `project/PHASE_STATUS.md` untuk status terkini.

```text
1. Pilih phase dengan status TODO
2. Baca file phase di /phases
3. Implement sesuai acceptance criteria
4. Test
5. Update PHASE_STATUS.md
6. Commit dengan format: feat(phase-XX): description
```

## License

Private — for internal use only.
