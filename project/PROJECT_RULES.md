# PROJECT_RULES.md

# Project Rules — Customer Communication Dashboard

## 1. Project Identity

Project ini adalah Customer Communication Dashboard untuk mengelola pesan customer dari WhatsApp dan Telegram dalam satu dashboard.

Stack utama:

* Backend: Golang + Gin
* Frontend: React + Vite + TypeScript
* Database: PostgreSQL
* Queue/Cache/Realtime: Redis
* Object Storage: MinIO
* Container: Docker + Docker Compose
* Orchestration: Kubernetes
* Integration: Telegram Bot API dan WhatsApp API/Gateway

MVP fokus pada:

* Unified inbox
* Reply manual
* Auto-reply keyword
* Template pesan
* Histori percakapan
* Attachment via MinIO
* Queue Redis
* Webhook Telegram
* Webhook WhatsApp
* Realtime inbox notification

---

## 2. Development Principle

AI coding agent harus bekerja secara bertahap.

Aturan utama:

1. Kerjakan project berdasarkan phase.
2. Jangan loncat phase.
3. Jangan membuat fitur yang belum masuk phase saat ini.
4. Setiap phase harus punya acceptance criteria.
5. Setiap phase harus dites sebelum lanjut.
6. Jika test gagal, perbaiki dulu.
7. Jika phase sukses, update `PHASE_STATUS.md`.
8. Jangan hardcode secret, token, password, atau API key.
9. Gunakan environment variable.
10. Gunakan struktur folder modular.
11. Backend, frontend, worker, dan infrastructure harus dipisahkan dengan jelas.
12. Jangan membuat AI chatbot sebelum MVP selesai.
13. Jangan membuat broadcast campaign sebelum MVP selesai.
14. Jangan membuat payment gateway sebelum MVP selesai.
15. Jangan membuat multi-tenant sebelum MVP selesai.

---

## 3. Phase Workflow

Setiap phase wajib mengikuti alur ini:

```text
Read PROJECT_RULES.md
  ↓
Read PHASE_STATUS.md
  ↓
Find first TODO phase
  ↓
Read phase file in /phases
  ↓
Implement only current phase
  ↓
Run test command
  ↓
Fix errors if any
  ↓
Update PHASE_STATUS.md
  ↓
Write phase report
  ↓
Stop and wait for next instruction
```

---

## 4. Status Rules

Status phase yang digunakan:

```text
TODO
IN_PROGRESS
PASSED
FAILED
SKIPPED
```

Aturan update status:

* Saat mulai mengerjakan phase, ubah status menjadi `IN_PROGRESS`.
* Jika semua acceptance criteria berhasil, ubah status menjadi `PASSED`.
* Jika masih ada error, ubah status menjadi `FAILED`.
* Jangan ubah phase berikutnya menjadi `IN_PROGRESS` sebelum phase saat ini `PASSED`.

---

## 5. Required Report After Each Phase

Setelah menyelesaikan phase, AI wajib membuat laporan:

```text
Phase:
Status:
Summary:
Files created:
Files changed:
Commands executed:
Test result:
Known issues:
Next phase:
```

Contoh:

```text
Phase: 01 — Project Setup
Status: PASSED

Summary:
Backend Golang berhasil dibuat dengan health check endpoint, koneksi PostgreSQL, Redis, dan MinIO.

Files created:
- backend/cmd/api/main.go
- backend/internal/config/config.go
- backend/internal/database/postgres.go
- backend/internal/database/redis.go
- backend/internal/database/minio.go
- backend/Dockerfile
- docker-compose.yml

Commands executed:
- docker compose up -d
- curl http://localhost:8080/health

Test result:
GET /health returned 200 OK.

Known issues:
None.

Next phase:
Phase 02 — Authentication
```

---

## 6. Backend Rules

Backend menggunakan Golang.

Framework utama:

* Gin

Backend harus menggunakan struktur modular:

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

Aturan backend:

1. Setiap module punya handler, service, repository, model, dan DTO jika diperlukan.
2. Handler hanya menangani HTTP request/response.
3. Service berisi business logic.
4. Repository hanya untuk database query.
5. Provider external seperti Telegram dan WhatsApp harus dipisahkan.
6. Redis queue logic harus dipisahkan dari handler.
7. MinIO logic harus dipisahkan ke storage service.
8. Jangan mencampur logic webhook, database, dan provider dalam satu file besar.
9. Gunakan context timeout untuk request external.
10. Gunakan response JSON yang konsisten.

Format response sukses:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

Format response error:

```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error"
}
```

---

## 7. Frontend Rules

Frontend menggunakan React.

Stack frontend:

* React
* Vite
* TypeScript
* Tailwind CSS
* React Router
* TanStack Query
* Zustand
* Axios

Struktur frontend:

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

Aturan frontend:

1. Semua API call harus melalui `services/api.ts`.
2. Auth state disimpan di Zustand.
3. Server state menggunakan TanStack Query.
4. Jangan hardcode URL API.
5. Gunakan environment variable `VITE_API_URL`.
6. Setiap halaman harus punya loading state.
7. Setiap halaman harus punya error state.
8. Setiap list harus punya empty state.
9. UI harus sederhana, clean, dan responsive.
10. Jangan membuat UI terlalu kompleks untuk MVP.

---

## 8. Database Rules

Database utama adalah PostgreSQL.

Aturan database:

1. Gunakan UUID sebagai primary key.
2. Semua table utama wajib punya `created_at`.
3. Table yang bisa di-update wajib punya `updated_at`.
4. Gunakan foreign key.
5. Tambahkan index untuk kolom yang sering difilter.
6. Jangan menyimpan file fisik di PostgreSQL.
7. File fisik harus disimpan di MinIO.
8. PostgreSQL hanya menyimpan metadata attachment.

Table utama:

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

---

## 9. Redis Rules

Redis digunakan untuk:

* Queue pesan keluar
* Retry pesan gagal
* Idempotency webhook
* Pub/Sub realtime notification
* Cache dashboard summary
* Rate limiting

Redis tidak boleh digunakan sebagai database utama.

Queue keys:

```text
queue:send-message
queue:send-telegram
queue:send-whatsapp
queue:retry-message
```

Pub/Sub channels:

```text
pubsub:inbox:new-message
pubsub:conversation:updated
pubsub:message:status-updated
```

Idempotency keys:

```text
idempotency:telegram:{provider_message_id}
idempotency:whatsapp:{provider_message_id}
```

Cache keys:

```text
cache:dashboard:summary:{date}
cache:conversation:unread_count
cache:channel:whatsapp:today
cache:channel:telegram:today
```

Aturan Redis:

1. Semua webhook harus cek idempotency.
2. Pesan keluar harus masuk queue dulu.
3. Worker yang mengirim pesan ke provider.
4. Jika gagal, retry maksimal 5 kali.
5. Jika tetap gagal, status message menjadi `failed`.
6. Event realtime harus dipublish setelah pesan baru atau status berubah.

---

## 10. MinIO Rules

MinIO digunakan untuk menyimpan attachment.

Jenis file:

* Gambar
* Dokumen
* PDF
* Voice note
* Bukti transfer
* File komplain
* File katalog

Bucket default:

```text
chat-media
```

Format object key:

```text
chat-media/{year}/{month}/{conversation_id}/{uuid}-{filename}
```

Aturan MinIO:

1. Upload file menggunakan presigned URL.
2. Download/view file menggunakan signed URL.
3. Jangan membuat bucket public untuk MVP.
4. Validasi ukuran file.
5. Validasi MIME type.
6. Simpan metadata file ke table `message_attachments`.

---

## 11. Webhook Rules

Webhook wajib:

1. Validasi payload.
2. Cek idempotency Redis.
3. Parse customer identifier.
4. Cari atau buat customer.
5. Cari atau buat conversation.
6. Simpan message.
7. Jika ada attachment, simpan ke MinIO.
8. Publish realtime event.
9. Cek auto-reply rule.
10. Jika auto-reply cocok, push job ke Redis queue.

Webhook endpoint:

```text
POST /api/webhooks/telegram
POST /api/webhooks/whatsapp
```

Webhook tidak boleh langsung mengirim banyak proses berat dalam request utama.

---

## 12. Worker Rules

Worker bertugas memproses queue Redis.

Worker harus:

1. Ambil job dari Redis queue.
2. Cek message ID di PostgreSQL.
3. Kirim pesan sesuai channel.
4. Update status message.
5. Retry jika gagal.
6. Publish status update ke Redis Pub/Sub.
7. Log error dengan jelas.

Worker tidak boleh menerima HTTP request.

---

## 13. Docker Rules

Local development wajib bisa dijalankan dengan:

```bash
docker compose up -d
```

Service minimal:

```text
postgres
redis
minio
backend
frontend
```

Aturan Docker:

1. Jangan hardcode environment di Dockerfile.
2. Gunakan `.env`.
3. Gunakan volume untuk PostgreSQL, Redis, dan MinIO.
4. Backend expose port 8080.
5. Frontend expose port 5173.
6. MinIO console expose port 9001.

---

## 14. Kubernetes Rules

Kubernetes digunakan untuk production atau staging.

Minimal object:

```text
Namespace
ConfigMap
Secret
Deployment backend
Deployment worker
Deployment frontend
Service backend
Service frontend
Ingress
PVC PostgreSQL
PVC Redis
PVC MinIO
```

Aturan Kubernetes:

1. Secret tidak boleh ditulis plain text di dokumentasi production.
2. Gunakan ConfigMap untuk config non-secret.
3. Gunakan Secret untuk token/password.
4. Gunakan PVC untuk storage.
5. Backend dan worker harus deployment terpisah.
6. Frontend dan backend harus service terpisah.
7. Ingress digunakan untuk domain.

---

## 15. Security Rules

Wajib:

1. Password user di-hash menggunakan bcrypt.
2. Auth menggunakan JWT.
3. Secret dari environment variable.
4. API protected menggunakan middleware.
5. Webhook memakai verify token atau secret jika provider mendukung.
6. File download menggunakan signed URL.
7. Rate limit login endpoint.
8. Rate limit webhook endpoint.
9. Jangan log token, password, atau secret.
10. Jangan commit `.env`.

---

## 16. Testing Rules

Setiap phase harus punya test command.

Minimal test:

```bash
docker compose up -d
curl http://localhost:8080/health
```

Untuk API:

```bash
curl -X POST http://localhost:8080/api/auth/login
curl http://localhost:8080/api/conversations
```

Untuk frontend:

```bash
npm install
npm run dev
npm run build
```

Untuk backend:

```bash
go test ./...
go run ./cmd/api
```

Untuk worker:

```bash
go run ./cmd/worker
```

---

## 17. Git Rules

Disarankan setiap phase selesai langsung commit.

Format commit:

```text
feat(phase-01): setup backend and docker compose
feat(phase-02): add JWT authentication
feat(phase-03): add core inbox module
feat(phase-04): add Telegram webhook
feat(phase-05): add Redis queue worker
```

Jangan commit:

```text
.env
node_modules
tmp files
MinIO data
PostgreSQL data
Redis data
```

---

## 18. DeepSeek Instruction

Saat menggunakan DeepSeek, selalu mulai dengan instruksi:

```text
Baca PROJECT_RULES.md dan PHASE_STATUS.md terlebih dahulu.
Kerjakan hanya phase pertama yang statusnya TODO.
Jangan lanjut ke phase berikutnya sebelum phase saat ini PASSED.
Ikuti acceptance criteria di file phase.
Setelah selesai, update PHASE_STATUS.md dan buat laporan phase.
```

---

## 19. Stop Condition

AI harus berhenti setelah menyelesaikan satu phase.

AI tidak boleh otomatis lanjut ke phase berikutnya kecuali user memberi instruksi:

```text
Lanjut phase berikutnya.
```

Tujuannya agar user bisa review kode, test manual, dan commit sebelum fitur berikutnya dibuat.
