# SECURITY_CHECKLIST.md

# Security Checklist — Customer Communication Dashboard

## 1. Authentication

* [ ] Password is hashed using bcrypt.
* [ ] Plain password is never stored.
* [ ] JWT secret comes from environment variable.
* [ ] JWT token has expiration time.
* [ ] Protected API routes use auth middleware.
* [ ] Admin-only routes validate user role.
* [ ] Login endpoint has rate limiting.
* [ ] Invalid login returns generic error.
* [ ] Logout invalidates token if blacklist is implemented.
* [ ] Password is never returned in API response.

---

## 2. Environment and Secrets

* [ ] `.env` is not committed.
* [ ] `.env.example` contains dummy values only.
* [ ] Telegram token is not hardcoded.
* [ ] WhatsApp token is not hardcoded.
* [ ] MinIO secret key is not hardcoded.
* [ ] PostgreSQL password is not hardcoded.
* [ ] JWT secret is not hardcoded.
* [ ] Kubernetes uses Secret for sensitive values.
* [ ] ConfigMap only stores non-sensitive values.
* [ ] Logs do not print secrets.

---

## 3. Webhook Security

* [ ] Telegram webhook uses secret token if supported.
* [ ] WhatsApp webhook uses verify token.
* [ ] Webhook payload is validated.
* [ ] Webhook request size is limited.
* [ ] Duplicate webhook is blocked using Redis idempotency.
* [ ] Provider message ID is stored.
* [ ] Webhook endpoint has rate limiting.
* [ ] Invalid webhook payload returns safe error.
* [ ] Webhook does not run heavy process synchronously.
* [ ] Webhook errors do not crash backend.

---

## 4. Redis Security

* [ ] Redis is not exposed publicly.
* [ ] Redis is only accessible inside Docker/Kubernetes network.
* [ ] Redis password is used in production if exposed internally.
* [ ] Redis does not store permanent business data.
* [ ] Redis keys have TTL where needed.
* [ ] Idempotency keys have expiration.
* [ ] Cache keys have expiration.
* [ ] Queue worker handles malformed jobs safely.
* [ ] Retry is limited to avoid infinite loop.
* [ ] Failed jobs are logged.

---

## 5. MinIO Security

* [ ] MinIO console is not publicly exposed in production.
* [ ] Bucket is not public.
* [ ] Upload uses presigned URL.
* [ ] Download uses signed URL.
* [ ] Signed URL has expiration.
* [ ] File size is validated.
* [ ] MIME type is validated.
* [ ] Dangerous file types are rejected.
* [ ] Original filename is sanitized.
* [ ] Object key does not expose sensitive information.

Allowed MVP MIME types:

```text
image/jpeg
image/png
image/webp
application/pdf
audio/mpeg
audio/ogg
video/mp4
```

Blocked file types:

```text
.exe
.sh
.bat
.php
.js
.html
```

---

## 6. API Security

* [ ] All protected endpoints require JWT.
* [ ] API validates request body.
* [ ] API validates UUID parameters.
* [ ] API validates enum values.
* [ ] API uses consistent error response.
* [ ] API does not leak stack trace.
* [ ] API uses CORS whitelist.
* [ ] API has request timeout.
* [ ] API has body size limit.
* [ ] API uses HTTPS in production.

---

## 7. Frontend Security

* [ ] API URL comes from environment variable.
* [ ] Token is not printed to console.
* [ ] Protected pages check auth state.
* [ ] User input is escaped.
* [ ] Chat message content is not rendered as raw HTML.
* [ ] Attachment preview validates file type.
* [ ] Logout clears auth state.
* [ ] Error messages do not expose secret.
* [ ] No production secret in frontend code.
* [ ] Build does not include debug-only secret.

---

## 8. Database Security

* [ ] Database is not exposed publicly.
* [ ] Database password is strong.
* [ ] Queries use parameterized statements.
* [ ] Foreign keys are enabled.
* [ ] Sensitive provider settings are encrypted if stored.
* [ ] Indexes exist for frequently queried fields.
* [ ] User role is validated.
* [ ] Deleted records are handled intentionally.
* [ ] Migration files are versioned.
* [ ] Backup is encrypted if stored externally.

---

## 9. Docker Security

* [ ] `.env` is not copied into image.
* [ ] Docker image uses minimal base image where possible.
* [ ] Container does not run unnecessary services.
* [ ] Volumes are used for persistent data.
* [ ] Production ports are limited.
* [ ] Redis and PostgreSQL are not publicly exposed.
* [ ] MinIO console is protected.
* [ ] Backend and worker use separate containers.
* [ ] Docker logs do not expose tokens.
* [ ] Unused containers/images are cleaned periodically.

---

## 10. Kubernetes Security

* [ ] Secrets use Kubernetes Secret.
* [ ] ConfigMap does not contain passwords.
* [ ] Ingress uses TLS.
* [ ] Backend and frontend have separate services.
* [ ] Worker has no public service.
* [ ] PostgreSQL, Redis, and MinIO are internal services.
* [ ] PVC is configured correctly.
* [ ] Resource limits are configured.
* [ ] Readiness and liveness probes are configured.
* [ ] ServiceAccount has minimal permission.

---

## 11. Logging and Monitoring

* [ ] Error logs include request ID.
* [ ] Logs do not include password.
* [ ] Logs do not include JWT token.
* [ ] Logs do not include provider access token.
* [ ] Webhook failures are logged.
* [ ] Queue failures are logged.
* [ ] Failed messages are traceable.
* [ ] Backend panic is recovered.
* [ ] Health check endpoint exists.
* [ ] Worker has health or log monitoring.

---

## 12. Final Security Acceptance

MVP is secure enough for initial deployment if:

* [ ] No hardcoded secrets.
* [ ] Auth works.
* [ ] Protected routes are protected.
* [ ] Webhook idempotency works.
* [ ] File upload is validated.
* [ ] MinIO bucket is private.
* [ ] Redis is internal only.
* [ ] PostgreSQL is internal only.
* [ ] HTTPS is enabled.
* [ ] Logs do not leak sensitive data.

---

# ARCHITECTURE.md

# Architecture — Customer Communication Dashboard

## 1. System Overview

Customer Communication Dashboard adalah aplikasi untuk mengelola pesan customer dari WhatsApp dan Telegram dalam satu dashboard.

MVP berfokus pada:

* Unified inbox
* Manual reply
* Keyword auto-reply
* Message template
* Conversation history
* Attachment storage
* Redis queue
* Realtime inbox notification

---

## 2. Technology Stack

```text
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
```

---

## 3. High Level Architecture

```text
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
```

---

## 4. Component Architecture

```text
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
```

---

## 5. Backend Architecture

Backend uses modular architecture.

```text
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
```

Layer responsibility:

```text
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
```

---

## 6. Frontend Architecture

```text
frontend/
  src/
    app/
    pages/
    components/
    services/
    stores/
    types/
```

Frontend responsibility:

```text
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
```

---

## 7. Data Storage Architecture

### PostgreSQL

Stores structured business data:

```text
users
customers
conversations
messages
message_attachments
message_templates
auto_reply_rules
provider_settings
```

### Redis

Used for:

```text
queue
retry
idempotency
pub/sub
cache
rate limit
```

### MinIO

Used for:

```text
image attachment
PDF attachment
voice note
video
customer document
proof of payment
```

---

## 8. Message Flow — Incoming Message

```text
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
```

---

## 9. Message Flow — Admin Reply

```text
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
```

---

## 10. Attachment Flow

```text
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
```

---

## 11. Redis Queue Architecture

Queue names:

```text
queue:send-message
queue:send-telegram
queue:send-whatsapp
queue:retry-message
```

Job payload:

```json
{
  "message_id": "uuid",
  "conversation_id": "uuid",
  "channel": "telegram",
  "body": "Hello",
  "attachment_id": null,
  "attempt": 1
}
```

Retry strategy:

```text
Attempt 1: immediate
Attempt 2: 30 seconds
Attempt 3: 2 minutes
Attempt 4: 5 minutes
Attempt 5: failed
```

---

## 12. Realtime Architecture

MVP uses Server-Sent Events.

```text
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
```

Events:

```text
new-message
conversation-updated
message-status-updated
```

---

## 13. Deployment Architecture

### Local

```text
Docker Compose
  |-- backend
  |-- worker
  |-- frontend
  |-- postgres
  |-- redis
  |-- minio
```

### Production MVP

```text
VPS
  |-- Nginx
  |-- Docker Compose
      |-- backend
      |-- worker
      |-- frontend
      |-- postgres
      |-- redis
      |-- minio
```

### Scalable Production

```text
Kubernetes
  |-- Ingress
  |-- Frontend Deployment
  |-- Backend Deployment
  |-- Worker Deployment
  |-- PostgreSQL / Managed DB
  |-- Redis / Managed Redis
  |-- MinIO / S3-compatible storage
```

---

## 14. Security Architecture

Security principles:

```text
JWT for authentication
bcrypt for password
environment variable for secrets
private MinIO bucket
signed URL for file access
Redis idempotency for webhook duplicate prevention
rate limit for login and webhook
HTTPS in production
```

---

## 15. MVP Architecture Decision

For MVP, use:

```text
Golang + Gin
React + Vite
PostgreSQL
Redis
MinIO
Docker Compose
SSE
```

Kubernetes is prepared but should be implemented after MVP is stable.
