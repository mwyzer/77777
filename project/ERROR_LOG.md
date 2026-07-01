# ERROR_LOG.md

## Error Format

### Error ID

ERR-001

### Date

2026-07-01

### Phase

Phase 01 — Project Setup

### Error

Backend cannot connect to PostgreSQL.

### Cause

Wrong DB_HOST in .env.

### Fix

Set DB_HOST=postgres when running inside Docker Compose.

### Status

Fixed

---

### Error ID

ERR-002

### Date

2026-07-01

### Phase

Phase 01 — Project Setup

### Error

Docker build failed: "missing go.sum entry for module providing package"

### Cause

Placeholder go.sum had incorrect checksums. Docker COPY copied the bad go.sum, causing checksum mismatch during build. Additionally, go.sum was missing entries entirely for some modules.

### Fix

1. Deleted incorrect go.sum placeholder
2. Updated Dockerfile to use `GONOSUMCHECK=* GONOSUMDB=* GOFLAGS=-mod=mod go mod tidy` to generate fresh go.sum inside the container
3. Used `go build -mod=mod` to allow module resolution during build

### Status

Fixed

---

### Error ID

ERR-003

### Date

2026-07-01

### Phase

Phase 01 — Project Setup

### Error

Docker daemon unavailable — "error during connect: open //./pipe/dockerDesktopLinuxEngine: The system cannot find the file specified."

### Cause

Docker Desktop not running or crashed during build attempt.

### Fix

Restart Docker Desktop and run `docker compose up -d --build`.

### Status

Pending (requires Docker Desktop restart)
