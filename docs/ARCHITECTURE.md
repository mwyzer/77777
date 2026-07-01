Architecture — Customer Communication Dashboard
1. System Overview

Customer Communication Dashboard adalah aplikasi untuk mengelola pesan customer dari WhatsApp dan Telegram dalam satu dashboard.

MVP berfokus pada:

Unified inbox
Manual reply
Keyword auto-reply
Message template
Conversation history
Attachment storage
Redis queue
Realtime inbox notification
2. Technology Stack
Frontend:
- React
- Vite
- TypeScript
- Tailwind CSS
- React Router
- TanStack Query
- Zustand
- Axios

Backend:
- Golang
- Gin
- JWT Authentication
- PostgreSQL
- Redis
- MinIO

Infrastructure:
- Docker
- Docker Compose
- Kubernetes
- Nginx / Ingress

External Integration:
- Telegram Bot API
- WhatsApp Cloud API or WhatsApp Gateway
3. High Level Architecture
Customer
  |
  | WhatsApp / Telegram
  v
Webhook Provider
  |
  v
Golang Backend API
  |
  |-- PostgreSQL: main database
  |-- Redis: queue, cache, pub/sub, idempotency
  |-- MinIO: attachment storage
  |
  v
React Dashboard
  |
  v
Admin / Customer Service
4. Component Architecture
React Frontend
  |
  v
Golang REST API
  |
  |-- Auth Module
  |-- Dashboard Module
  |-- Customer Module
  |-- Conversation Module
  |-- Message Module
  |-- Template Module
  |-- Auto-Reply Module
  |-- Attachment Module
  |-- Webhook Module
  |
  |-- Telegram Provider
  |-- WhatsApp Provider
  |-- Redis Queue
  |-- Redis Pub/Sub
  |-- MinIO Storage
  |
  v
PostgreSQL / Redis / MinIO
5. Backend Architecture

Backend uses modular architecture.

backend/
  cmd/
    api/
      main.go
    worker/
      main.go

  internal/
    config/
    database/
    middleware/
    modules/
    providers/
    queue/
    realtime/
    storage/

Layer responsibility:

Handler:
- Receives HTTP request
- Validates request
- Returns JSON response

Service:
- Business logic
- Calls repository/provider/queue/storage

Repository:
- Database query only

Provider:
- External API integration
- Telegram
- WhatsApp

Queue:
- Redis queue
- Retry logic

Storage:
- MinIO upload/download logic
6. Frontend Architecture
frontend/
  src/
    app/
    pages/
    components/
    services/
    stores/
    types/

Frontend responsibility:

Pages:
- Route-level screens

Components:
- Reusable UI elements

Services:
- API client

Stores:
- Client state with Zustand

TanStack Query:
- Server state, cache, loading, refetch
7. Data Storage Architecture
PostgreSQL

Stores structured business data:

users
customers
conversations
messages
message_attachments
message_templates
auto_reply_rules
provider_settings
Redis

Used for:

queue
retry
idempotency
pub/sub
cache
rate limit
MinIO

Used for:

image attachment
PDF attachment
voice note
video
customer document
proof of payment
8. Message Flow — Incoming Message
Customer sends message
  |
  v
WhatsApp/Telegram provider
  |
  v
Webhook endpoint
  |
  v
Redis idempotency check
  |
  v
Create/find customer
  |
  v
Create/find conversation
  |
  v
Save message to PostgreSQL
  |
  v
Save attachment to MinIO if exists
  |
  v
Publish realtime event to Redis
  |
  v
Check auto-reply rule
  |
  v
Push auto-reply job to Redis queue if matched
9. Message Flow — Admin Reply
Admin writes reply in React
  |
  v
POST /api/conversations/:id/messages
  |
  v
Backend saves message as pending
  |
  v
Backend pushes job to Redis queue
  |
  v
Worker consumes job
  |
  v
Worker sends message to Telegram/WhatsApp
  |
  v
Worker updates message status
  |
  v
Worker publishes status event
  |
  v
React updates UI
10. Attachment Flow
Admin selects file
  |
  v
React requests presigned upload URL
  |
  v
Backend creates MinIO presigned URL
  |
  v
React uploads file directly to MinIO
  |
  v
React confirms upload to backend
  |
  v
Backend saves metadata to PostgreSQL
  |
  v
Message uses attachment ID
11. Redis Queue Architecture

Queue names:

queue:send-message
queue:send-telegram
queue:send-whatsapp
queue:retry-message

Job payload:

{
  "message_id": "uuid",
  "conversation_id": "uuid",
  "channel": "telegram",
  "body": "Hello",
  "attachment_id": null,
  "attempt": 1
}

Retry strategy:

Attempt 1: immediate
Attempt 2: 30 seconds
Attempt 3: 2 minutes
Attempt 4: 5 minutes
Attempt 5: failed
12. Realtime Architecture

MVP uses Server-Sent Events.

Redis Pub/Sub
  |
  v
Backend SSE endpoint
  |
  v
React EventSource client
  |
  v
Inbox refresh / toast notification

Events:

new-message
conversation-updated
message-status-updated
13. Deployment Architecture
Local
Docker Compose
  |-- backend
  |-- worker
  |-- frontend
  |-- postgres
  |-- redis
  |-- minio
Production MVP
VPS
  |-- Nginx
  |-- Docker Compose
      |-- backend
      |-- worker
      |-- frontend
      |-- postgres
      |-- redis
      |-- minio
Scalable Production
Kubernetes
  |-- Ingress
  |-- Frontend Deployment
  |-- Backend Deployment
  |-- Worker Deployment
  |-- PostgreSQL / Managed DB
  |-- Redis / Managed Redis
  |-- MinIO / S3-compatible storage
14. Security Architecture

Security principles:

JWT for authentication
bcrypt for password
environment variable for secrets
private MinIO bucket
signed URL for file access
Redis idempotency for webhook duplicate prevention
rate limit for login and webhook
HTTPS in production
15. MVP Architecture Decision

For MVP, use:

Golang + Gin
React + Vite
PostgreSQL
Redis
MinIO
Docker Compose
SSE

Kubernetes is prepared but should be implemented after MVP is stable.