# Struktur Project Lengkap

## Customer Communication Dashboard

### Golang + React + PostgreSQL + Redis + MinIO + Docker + Kubernetes

```text
customer-communication-dashboard/
│
├── README.md
├── PROJECT_RULES.md
├── PHASE_STATUS.md
├── CONTEXT.md
├── TASKS.md
├── DECISIONS.md
├── ERROR_LOG.md
├── CHANGELOG.md
├── .gitignore
├── docker-compose.yml
│
├── docs/
│   ├── BRD.md
│   ├── SRS.md
│   ├── ARCHITECTURE.md
│   ├── API_CONTRACT.md
│   ├── DATABASE_SCHEMA.md
│   ├── DEPLOYMENT.md
│   ├── TEST_PLAN.md
│   └── SECURITY_CHECKLIST.md
│
├── phases/
│   ├── phase-01-project-setup.md
│   ├── phase-02-authentication.md
│   ├── phase-03-core-inbox-backend.md
│   ├── phase-04-react-dashboard-base.md
│   ├── phase-05-telegram-integration.md
│   ├── phase-06-redis-queue-worker.md
│   ├── phase-07-minio-attachment.md
│   ├── phase-08-whatsapp-integration.md
│   ├── phase-09-template-message.md
│   ├── phase-10-auto-reply-keyword.md
│   ├── phase-11-realtime-inbox.md
│   ├── phase-12-dashboard-summary.md
│   ├── phase-13-kubernetes-deployment.md
│   └── phase-14-final-review.md
│
├── prompts/
│   ├── master-prompt.md
│   ├── continue-next-phase.md
│   ├── fix-error.md
│   ├── review-code.md
│   └── refactor.md
│
├── scripts/
│   ├── check-health.sh
│   ├── check-phase-01.sh
│   ├── check-phase-02.sh
│   ├── run-backend-tests.sh
│   ├── run-frontend-tests.sh
│   ├── run-all-tests.sh
│   └── reset-local.sh
│
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   │   └── main.go
│   │   └── worker/
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   │
│   │   ├── database/
│   │   │   ├── postgres.go
│   │   │   ├── redis.go
│   │   │   └── minio.go
│   │   │
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go
│   │   │   ├── rate_limit_middleware.go
│   │   │   ├── cors_middleware.go
│   │   │   └── recovery_middleware.go
│   │   │
│   │   ├── response/
│   │   │   └── response.go
│   │   │
│   │   ├── modules/
│   │   │   ├── auth/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── users/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── customers/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── conversations/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── messages/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── attachments/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── templates/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── autoreply/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   ├── repository.go
│   │   │   │   ├── matcher.go
│   │   │   │   ├── dto.go
│   │   │   │   └── model.go
│   │   │   │
│   │   │   ├── dashboard/
│   │   │   │   ├── handler.go
│   │   │   │   ├── service.go
│   │   │   │   └── dto.go
│   │   │   │
│   │   │   └── webhooks/
│   │   │       ├── telegram_handler.go
│   │   │       ├── whatsapp_handler.go
│   │   │       ├── telegram_parser.go
│   │   │       └── whatsapp_parser.go
│   │   │
│   │   ├── providers/
│   │   │   ├── telegram/
│   │   │   │   ├── client.go
│   │   │   │   └── dto.go
│   │   │   │
│   │   │   └── whatsapp/
│   │   │       ├── provider.go
│   │   │       ├── cloud_api.go
│   │   │       ├── fonnte.go
│   │   │       └── dto.go
│   │   │
│   │   ├── queue/
│   │   │   ├── redis_queue.go
│   │   │   ├── jobs.go
│   │   │   ├── send_message_job.go
│   │   │   └── retry.go
│   │   │
│   │   ├── realtime/
│   │   │   ├── pubsub.go
│   │   │   └── sse.go
│   │   │
│   │   ├── storage/
│   │   │   └── minio_service.go
│   │   │
│   │   ├── idempotency/
│   │   │   └── redis_idempotency.go
│   │   │
│   │   └── validator/
│   │       └── validator.go
│   │
│   ├── migrations/
│   │   ├── 001_create_users.sql
│   │   ├── 002_create_customers.sql
│   │   ├── 003_create_conversations.sql
│   │   ├── 004_create_messages.sql
│   │   ├── 005_create_message_attachments.sql
│   │   ├── 006_create_message_templates.sql
│   │   ├── 007_create_auto_reply_rules.sql
│   │   ├── 008_create_provider_settings.sql
│   │   └── 009_seed_default_admin.sql
│   │
│   ├── tests/
│   │   ├── auth_test.go
│   │   ├── conversation_test.go
│   │   ├── message_test.go
│   │   ├── telegram_webhook_test.go
│   │   ├── redis_queue_test.go
│   │   └── minio_attachment_test.go
│   │
│   ├── .env.example
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   ├── router.tsx
│   │   │   ├── providers.tsx
│   │   │   └── protected-route.tsx
│   │   │
│   │   ├── pages/
│   │   │   ├── LoginPage.tsx
│   │   │   ├── DashboardPage.tsx
│   │   │   ├── InboxPage.tsx
│   │   │   ├── ConversationDetailPage.tsx
│   │   │   ├── TemplatesPage.tsx
│   │   │   ├── AutoReplyPage.tsx
│   │   │   ├── SettingsPage.tsx
│   │   │   └── NotFoundPage.tsx
│   │   │
│   │   ├── components/
│   │   │   ├── layout/
│   │   │   │   ├── Sidebar.tsx
│   │   │   │   ├── Navbar.tsx
│   │   │   │   └── DashboardLayout.tsx
│   │   │   │
│   │   │   ├── dashboard/
│   │   │   │   └── SummaryCard.tsx
│   │   │   │
│   │   │   ├── inbox/
│   │   │   │   ├── ConversationList.tsx
│   │   │   │   ├── ConversationItem.tsx
│   │   │   │   ├── ChatWindow.tsx
│   │   │   │   ├── MessageBubble.tsx
│   │   │   │   ├── ReplyBox.tsx
│   │   │   │   ├── TemplatePicker.tsx
│   │   │   │   ├── AttachmentPreview.tsx
│   │   │   │   └── StatusBadge.tsx
│   │   │   │
│   │   │   └── common/
│   │   │       ├── Button.tsx
│   │   │       ├── Input.tsx
│   │   │       ├── LoadingState.tsx
│   │   │       ├── ErrorState.tsx
│   │   │       └── EmptyState.tsx
│   │   │
│   │   ├── services/
│   │   │   ├── api.ts
│   │   │   ├── authApi.ts
│   │   │   ├── dashboardApi.ts
│   │   │   ├── conversationApi.ts
│   │   │   ├── messageApi.ts
│   │   │   ├── attachmentApi.ts
│   │   │   ├── templateApi.ts
│   │   │   └── autoReplyApi.ts
│   │   │
│   │   ├── stores/
│   │   │   ├── authStore.ts
│   │   │   └── inboxStore.ts
│   │   │
│   │   ├── hooks/
│   │   │   ├── useAuth.ts
│   │   │   ├── useConversations.ts
│   │   │   ├── useMessages.ts
│   │   │   └── useRealtime.ts
│   │   │
│   │   ├── types/
│   │   │   ├── auth.ts
│   │   │   ├── dashboard.ts
│   │   │   ├── customer.ts
│   │   │   ├── conversation.ts
│   │   │   ├── message.ts
│   │   │   ├── attachment.ts
│   │   │   ├── template.ts
│   │   │   └── autoReply.ts
│   │   │
│   │   ├── utils/
│   │   │   ├── formatDate.ts
│   │   │   ├── cn.ts
│   │   │   └── constants.ts
│   │   │
│   │   ├── main.tsx
│   │   └── index.css
│   │
│   ├── public/
│   ├── .env.example
│   ├── Dockerfile
│   ├── index.html
│   ├── package.json
│   ├── package-lock.json
│   ├── tsconfig.json
│   ├── tailwind.config.js
│   └── vite.config.ts
│
├── k8s/
│   ├── namespace.yaml
│   ├── configmap.yaml
│   ├── secret.example.yaml
│   ├── backend-deployment.yaml
│   ├── backend-service.yaml
│   ├── worker-deployment.yaml
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   ├── postgres-deployment.yaml
│   ├── postgres-service.yaml
│   ├── postgres-pvc.yaml
│   ├── redis-deployment.yaml
│   ├── redis-service.yaml
│   ├── redis-pvc.yaml
│   ├── minio-deployment.yaml
│   ├── minio-service.yaml
│   ├── minio-pvc.yaml
│   └── ingress.yaml
│
└── postman/
    └── customer-communication-dashboard.postman_collection.json
```

---

# Fungsi Setiap Bagian

## 1. Root Files

```text
README.md
```

Penjelasan umum project, cara install, cara run, dan ringkasan fitur.

```text
PROJECT_RULES.md
```

Aturan utama untuk DeepSeek atau AI coding agent agar tidak loncat phase.

```text
PHASE_STATUS.md
```

Status progress setiap phase: TODO, IN_PROGRESS, PASSED, FAILED.

```text
CONTEXT.md
```

Konteks singkat project agar AI tidak lupa tujuan aplikasi.

```text
TASKS.md
```

Checklist teknis per phase.

```text
DECISIONS.md
```

Catatan keputusan teknis, misalnya kenapa pakai Gin, SSE, Redis, dan MinIO.

```text
ERROR_LOG.md
```

Catatan error, root cause, dan solusi.

```text
CHANGELOG.md
```

Riwayat perubahan project.

---

# 2. Folder `docs/`

Folder ini berisi dokumen utama project.

```text
BRD.md
```

Business Requirement Document. Menjelaskan kebutuhan bisnis.

```text
SRS.md
```

Software Requirement Specification. Menjelaskan kebutuhan sistem secara teknis.

```text
ARCHITECTURE.md
```

Desain arsitektur sistem.

```text
API_CONTRACT.md
```

Kontrak API antara backend dan frontend.

```text
DATABASE_SCHEMA.md
```

Struktur tabel database.

```text
DEPLOYMENT.md
```

Panduan deployment Docker Compose dan Kubernetes.

```text
TEST_PLAN.md
```

Rencana testing per phase.

```text
SECURITY_CHECKLIST.md
```

Checklist keamanan.

---

# 3. Folder `phases/`

Folder ini penting untuk vibe coding. DeepSeek membaca satu phase, mengerjakan satu tahap, lalu berhenti.

Urutan phase:

```text
phase-01-project-setup.md
phase-02-authentication.md
phase-03-core-inbox-backend.md
phase-04-react-dashboard-base.md
phase-05-telegram-integration.md
phase-06-redis-queue-worker.md
phase-07-minio-attachment.md
phase-08-whatsapp-integration.md
phase-09-template-message.md
phase-10-auto-reply-keyword.md
phase-11-realtime-inbox.md
phase-12-dashboard-summary.md
phase-13-kubernetes-deployment.md
phase-14-final-review.md
```

Aturan terbaik:

```text
1 phase = 1 fokus = 1 test = 1 commit
```

---

# 4. Folder `prompts/`

Folder ini berisi prompt siap pakai untuk DeepSeek.

```text
master-prompt.md
```

Prompt utama untuk memulai project.

```text
continue-next-phase.md
```

Prompt untuk lanjut phase berikutnya.

```text
fix-error.md
```

Prompt khusus memperbaiki error.

```text
review-code.md
```

Prompt untuk review kode.

```text
refactor.md
```

Prompt untuk refactor aman tanpa mengubah fitur.

---

# 5. Folder `scripts/`

Berisi script bantu agar testing tidak manual terus.

```text
check-health.sh
```

Cek endpoint health backend.

```text
check-phase-01.sh
```

Cek hasil Phase 01.

```text
check-phase-02.sh
```

Cek hasil Phase 02.

```text
run-backend-tests.sh
```

Menjalankan test backend.

```text
run-frontend-tests.sh
```

Menjalankan build/test frontend.

```text
run-all-tests.sh
```

Menjalankan semua test.

```text
reset-local.sh
```

Reset environment lokal.

---

# 6. Folder `backend/`

Backend menggunakan Golang + Gin.

## Struktur Utama

```text
cmd/api/main.go
```

Entry point untuk REST API.

```text
cmd/worker/main.go
```

Entry point untuk worker Redis queue.

```text
internal/config/
```

Load environment variable.

```text
internal/database/
```

Koneksi PostgreSQL, Redis, dan MinIO.

```text
internal/middleware/
```

Auth middleware, rate limit, CORS, recovery.

```text
internal/modules/
```

Semua fitur utama aplikasi.

```text
internal/providers/
```

Integrasi Telegram dan WhatsApp.

```text
internal/queue/
```

Redis queue dan retry logic.

```text
internal/realtime/
```

Redis Pub/Sub dan SSE.

```text
internal/storage/
```

Service MinIO.

```text
migrations/
```

SQL migration.

---

# 7. Folder `frontend/`

Frontend menggunakan React + Vite + TypeScript.

## Struktur Utama

```text
src/pages/
```

Halaman utama aplikasi.

```text
src/components/
```

Komponen UI.

```text
src/services/
```

API client.

```text
src/stores/
```

State management Zustand.

```text
src/hooks/
```

Custom React hooks.

```text
src/types/
```

TypeScript type definitions.

```text
src/utils/
```

Helper function.

---

# 8. Folder `k8s/`

Folder ini berisi manifest Kubernetes.

Minimal komponen:

```text
namespace
configmap
secret
backend deployment
worker deployment
frontend deployment
postgres deployment
redis deployment
minio deployment
service
pvc
ingress
```

Catatan:

Untuk MVP awal, jalankan pakai Docker Compose dulu. Kubernetes dipakai setelah fitur utama stabil.

---

# 9. Folder `postman/`

Berisi collection untuk testing API.

Contoh isi:

```text
Auth
Dashboard
Conversations
Messages
Attachments
Templates
Auto Reply
Telegram Webhook
WhatsApp Webhook
```

---

# Struktur Minimum Kalau Mau Lebih Simpel

Kalau ingin tidak terlalu berat di awal, gunakan struktur minimum ini dulu:

```text
customer-communication-dashboard/
│
├── README.md
├── PROJECT_RULES.md
├── PHASE_STATUS.md
├── CONTEXT.md
├── TASKS.md
├── docker-compose.yml
│
├── docs/
│   ├── BRD.md
│   ├── SRS.md
│   ├── ARCHITECTURE.md
│   ├── API_CONTRACT.md
│   ├── DATABASE_SCHEMA.md
│   ├── TEST_PLAN.md
│   └── SECURITY_CHECKLIST.md
│
├── phases/
│   ├── phase-01-project-setup.md
│   ├── phase-02-authentication.md
│   ├── phase-03-core-inbox-backend.md
│   └── phase-04-react-dashboard-base.md
│
├── prompts/
│   ├── master-prompt.md
│   ├── continue-next-phase.md
│   └── fix-error.md
│
├── backend/
├── frontend/
└── k8s/
```

Setelah Phase 04 stabil, baru tambahkan file dan folder lain.

---

# Rekomendasi Final

Gunakan struktur lengkap, tapi kerjakan secara bertahap.

Urutan terbaik:

```text
1. Root documentation
2. docs/
3. phases/
4. prompts/
5. backend/
6. frontend/
7. docker-compose.yml
8. scripts/
9. k8s/
10. postman/
```

Untuk vibe coding dengan DeepSeek, file paling penting adalah:

```text
PROJECT_RULES.md
PHASE_STATUS.md
CONTEXT.md
TASKS.md
docs/ARCHITECTURE.md
docs/API_CONTRACT.md
docs/DATABASE_SCHEMA.md
docs/TEST_PLAN.md
phases/
prompts/
```


