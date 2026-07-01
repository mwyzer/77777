Master Prompt — Customer Communication Dashboard

You are a senior software engineer and system architect.

Build this project phase by phase.

Project name:

Customer Communication Dashboard

Description:

A dashboard to manage customer communication from WhatsApp and Telegram in one place.

Main MVP features:

- Unified inbox
- Manual reply
- Keyword auto-reply
- Message templates
- Conversation history
- Attachment storage with MinIO
- Redis queue
- Redis idempotency
- Redis Pub/Sub realtime notification
- Telegram webhook
- WhatsApp webhook
- Docker Compose
- Kubernetes basic manifest

Tech stack:

Backend:
- Golang
- Gin
- PostgreSQL
- Redis
- MinIO
- JWT Authentication

Frontend:
- React
- Vite
- TypeScript
- Tailwind CSS
- React Router
- TanStack Query
- Zustand
- Axios

Infrastructure:
- Docker
- Docker Compose
- Kubernetes

Read these files before coding:

PROJECT_RULES.md
PHASE_STATUS.md
CONTEXT.md
TASKS.md
docs/BRD.md
docs/SRS.md
docs/ARCHITECTURE.md
docs/API_CONTRACT.md
docs/DATABASE_SCHEMA.md
docs/TEST_PLAN.md
docs/SECURITY_CHECKLIST.md

Rules:

1. Work only on the first TODO phase.
2. Do not skip phases.
3. Do not implement future features early.
4. Follow acceptance criteria.
5. Run or explain test commands.
6. Fix errors before continuing.
7. Update PHASE_STATUS.md after finishing.
8. Update TASKS.md if task status changes.
9. Write a phase report.
10. Stop after one phase.

Do not build yet:

- AI chatbot
- Broadcast campaign
- Payment gateway
- Multi-tenant system
- Marketplace integration

Start with the first TODO phase.