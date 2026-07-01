# DECISIONS.md

## Decision 001 — Backend Framework

Use Gin for Golang backend.

Reason:
- Simple
- Fast
- Easy routing
- Good for REST API and webhook

## Decision 002 — Realtime Method

Use Server-Sent Events for MVP.

Reason:
- Simpler than WebSocket
- Enough for inbox notification
- Easier to implement with React

## Decision 003 — Storage

Use MinIO for file attachment.

Reason:
- S3-compatible
- Easy local development
- Can be replaced with S3 later

## Decision 004 — Redis Usage

Use Redis for:
- Queue
- Retry
- Idempotency
- Pub/Sub
- Cache

Redis is not used as primary database.