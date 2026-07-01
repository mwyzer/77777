# TEST_PLAN.md

# Test Plan — Customer Communication Dashboard

## 1. Project Overview

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

* Login admin
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
* Dashboard summary
* Docker Compose
* Kubernetes basic deployment

---

## 2. Testing Goals

Tujuan testing:

1. Memastikan backend berjalan dengan benar.
2. Memastikan frontend dapat terhubung ke backend.
3. Memastikan database PostgreSQL menyimpan data dengan benar.
4. Memastikan Redis berjalan untuk queue, cache, idempotency, dan pub/sub.
5. Memastikan MinIO dapat menyimpan dan mengakses attachment.
6. Memastikan webhook Telegram dan WhatsApp dapat menerima pesan.
7. Memastikan admin dapat membalas pesan dari dashboard.
8. Memastikan auto-reply keyword berjalan.
9. Memastikan sistem dapat dijalankan dengan Docker Compose.
10. Memastikan deployment Kubernetes basic dapat dijalankan.

---

## 3. General Testing Rules

Setiap phase harus memenuhi aturan berikut:

1. Semua command test harus dijalankan.
2. Tidak boleh lanjut phase berikutnya jika test utama gagal.
3. Error harus dicatat di `ERROR_LOG.md`.
4. Status phase harus diperbarui di `PHASE_STATUS.md`.
5. Semua API harus mengembalikan response JSON konsisten.
6. Tidak boleh ada secret/token/password yang hardcode.
7. Tidak boleh commit file `.env`.
8. Backend dan frontend harus bisa dijalankan ulang tanpa error fatal.
9. Docker Compose harus bisa dijalankan dari root project.
10. Setiap fitur harus punya minimal manual test.

---

## 4. Standard API Response Test

### Success Response Format

Semua response sukses harus mengikuti format:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

### Error Response Format

Semua response error harus mengikuti format:

```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error"
}
```

### Test Checklist

* [ ] Response sukses menggunakan `success: true`
* [ ] Response error menggunakan `success: false`
* [ ] Response memiliki field `message`
* [ ] Response memiliki field `data` atau `error`
* [ ] HTTP status code sesuai kondisi
* [ ] Error tidak membocorkan secret, token, atau password

---

# 5. Phase 01 — Project Setup Test

## Scope

Phase ini menguji setup awal project:

* Backend Golang
* Gin framework
* PostgreSQL connection
* Redis connection
* MinIO connection
* Docker Compose
* Health check endpoint

## Test Commands

```bash
docker compose up -d
```

```bash
docker compose ps
```

```bash
curl http://localhost:8080/health
```

```bash
docker compose logs backend
```

## Expected Result

* [ ] Semua container berjalan
* [ ] Backend berjalan di port `8080`
* [ ] PostgreSQL berjalan di port `5432`
* [ ] Redis berjalan di port `6379`
* [ ] MinIO berjalan di port `9000`
* [ ] MinIO Console berjalan di port `9001`
* [ ] Endpoint `/health` mengembalikan HTTP `200`
* [ ] Backend tidak menampilkan error fatal
* [ ] Backend berhasil connect ke PostgreSQL
* [ ] Backend berhasil connect ke Redis
* [ ] Backend berhasil connect ke MinIO

## Example Expected Health Response

```json
{
  "success": true,
  "message": "Service is healthy",
  "data": {
    "app": "ok",
    "postgres": "ok",
    "redis": "ok",
    "minio": "ok"
  }
}
```

## Failure Condition

Phase gagal jika:

* Container backend crash
* Endpoint `/health` tidak bisa diakses
* PostgreSQL gagal connect
* Redis gagal connect
* MinIO gagal connect
* Ada panic/error fatal di log backend

---

# 6. Phase 02 — Authentication Test

## Scope

Phase ini menguji:

* Login admin
* JWT token
* Password hashing
* Protected route
* Auth middleware
* Endpoint `/api/auth/me`

## Test Commands

### Login Admin

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password"}'
```

## Expected Result

* [ ] Login berhasil
* [ ] Response mengandung JWT token
* [ ] Response mengandung data user
* [ ] Password tidak muncul di response
* [ ] Token dapat digunakan untuk protected route

## Example Login Response

```json
{
  "success": true,
  "message": "Login success",
  "data": {
    "token": "jwt_token_here",
    "user": {
      "id": "uuid",
      "name": "Admin",
      "email": "admin@example.com",
      "role": "admin"
    }
  }
}
```

### Test Auth Me

```bash
curl http://localhost:8080/api/auth/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Endpoint mengembalikan data user
* [ ] Token valid diterima
* [ ] Token invalid ditolak
* [ ] Request tanpa token ditolak

### Test Without Token

```bash
curl http://localhost:8080/api/auth/me
```

## Expected Result

* [ ] HTTP status `401`
* [ ] Response error jelas
* [ ] Tidak ada data user dikembalikan

## Failure Condition

Phase gagal jika:

* Login gagal untuk default admin
* JWT token tidak dibuat
* Password muncul di response
* Protected route bisa diakses tanpa token
* Token invalid tetap diterima

---

# 7. Phase 03 — Core Inbox Backend Test

## Scope

Phase ini menguji:

* Customer
* Conversation
* Message
* Conversation list
* Conversation detail
* Message list
* Manual reply API
* Update status
* Update internal note

## Test Commands

### Get Conversations

```bash
curl http://localhost:8080/api/conversations \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Endpoint mengembalikan list conversation
* [ ] Response mendukung pagination
* [ ] Response memiliki channel
* [ ] Response memiliki status
* [ ] Response memiliki last message

### Get Conversation Detail

```bash
curl http://localhost:8080/api/conversations/CONVERSATION_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Detail conversation berhasil ditampilkan
* [ ] Customer data tampil
* [ ] Channel tampil
* [ ] Status tampil

### Get Messages

```bash
curl http://localhost:8080/api/conversations/CONVERSATION_ID/messages \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Histori pesan tampil
* [ ] Pesan terurut berdasarkan waktu
* [ ] Sender type tampil
* [ ] Delivery status tampil

### Create Manual Reply

```bash
curl -X POST http://localhost:8080/api/conversations/CONVERSATION_ID/messages \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body":"Halo kak, ada yang bisa kami bantu?","template_id":null,"attachment_id":null}'
```

## Expected Result

* [ ] Message baru dibuat
* [ ] Sender type adalah `admin`
* [ ] Delivery status awal adalah `pending`
* [ ] Message tersimpan di PostgreSQL

### Update Conversation Status

```bash
curl -X PATCH http://localhost:8080/api/conversations/CONVERSATION_ID/status \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"resolved"}'
```

## Expected Result

* [ ] Status conversation berubah
* [ ] Updated at berubah
* [ ] Status invalid ditolak

### Update Internal Note

```bash
curl -X PATCH http://localhost:8080/api/conversations/CONVERSATION_ID/note \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"internal_note":"Customer minta dikirim katalog terbaru."}'
```

## Expected Result

* [ ] Internal note berhasil disimpan
* [ ] Note tampil di detail conversation

## Failure Condition

Phase gagal jika:

* Conversation tidak bisa dilist
* Message tidak tersimpan
* Status invalid diterima
* Endpoint bisa diakses tanpa token
* Response tidak konsisten

---

# 8. Phase 04 — React Dashboard Base Test

## Scope

Phase ini menguji:

* React project
* Login page
* Dashboard layout
* Protected route
* Axios API client
* Zustand auth store
* Routing

## Test Commands

```bash
cd frontend
npm install
npm run dev
```

```bash
npm run build
```

## Expected Result

* [ ] Frontend berjalan di port `5173`
* [ ] Login page tampil
* [ ] User bisa login
* [ ] Token tersimpan di auth store
* [ ] Protected route tidak bisa diakses tanpa login
* [ ] Setelah login user diarahkan ke dashboard
* [ ] Build production berhasil

## Manual Browser Test

Buka:

```text
http://localhost:5173
```

Checklist:

* [ ] Login page tampil
* [ ] Input email tampil
* [ ] Input password tampil
* [ ] Tombol login tampil
* [ ] Login sukses masuk dashboard
* [ ] Sidebar tampil
* [ ] Navbar tampil
* [ ] Logout berjalan

## Failure Condition

Phase gagal jika:

* Frontend gagal dijalankan
* Build gagal
* Login tidak menyimpan token
* Protected route bisa diakses tanpa login
* API URL hardcode dan tidak memakai environment variable

---

# 9. Phase 05 — Telegram Integration Test

## Scope

Phase ini menguji:

* Telegram webhook
* Parse payload Telegram
* Redis idempotency
* Simpan customer
* Simpan conversation
* Simpan message
* Reply Telegram
* Publish realtime event

## Test Webhook Manually

```bash
curl -X POST http://localhost:8080/api/webhooks/telegram \
  -H "Content-Type: application/json" \
  -d '{
    "update_id": 100001,
    "message": {
      "message_id": 200001,
      "from": {
        "id": 123456789,
        "is_bot": false,
        "first_name": "Budi",
        "username": "budi_test"
      },
      "chat": {
        "id": 123456789,
        "first_name": "Budi",
        "username": "budi_test",
        "type": "private"
      },
      "date": 1710000000,
      "text": "Halo, saya mau tanya harga"
    }
  }'
```

## Expected Result

* [ ] Webhook menerima payload
* [ ] Redis menyimpan idempotency key
* [ ] Customer dibuat atau ditemukan
* [ ] Conversation channel `telegram` dibuat atau ditemukan
* [ ] Message tersimpan
* [ ] Provider message ID tersimpan
* [ ] Event realtime dipublish
* [ ] Response HTTP `200`

## Duplicate Webhook Test

Kirim payload yang sama dua kali.

Expected:

* [ ] Request pertama diproses
* [ ] Request kedua tidak membuat message duplikat
* [ ] Response tetap aman

## Telegram Reply Test

Dari dashboard atau API:

```bash
curl -X POST http://localhost:8080/api/conversations/CONVERSATION_ID/messages \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body":"Halo Budi, berikut info harga kami.","template_id":null,"attachment_id":null}'
```

Expected:

* [ ] Message masuk queue Redis
* [ ] Worker mengirim pesan ke Telegram
* [ ] Status berubah menjadi `sent` jika sukses
* [ ] Status berubah menjadi `failed` jika gagal

## Failure Condition

Phase gagal jika:

* Payload Telegram tidak bisa diparse
* Customer tidak tersimpan
* Message duplikat terjadi
* Reply tidak masuk queue
* Token Telegram hardcode
* Error provider menyebabkan backend crash

---

# 10. Phase 06 — Redis Queue Worker Test

## Scope

Phase ini menguji:

* Redis queue
* Worker
* Send message job
* Retry logic
* Status update
* Pub/Sub status event

## Test Commands

Jalankan API:

```bash
go run ./cmd/api
```

Jalankan worker:

```bash
go run ./cmd/worker
```

Atau via Docker Compose:

```bash
docker compose up -d backend worker redis
```

## Push Message Job Test

Buat pesan keluar dari API:

```bash
curl -X POST http://localhost:8080/api/conversations/CONVERSATION_ID/messages \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body":"Test queue message","template_id":null,"attachment_id":null}'
```

## Expected Result

* [ ] Message tersimpan dengan status `pending`
* [ ] Job masuk Redis queue
* [ ] Worker mengambil job
* [ ] Worker memproses berdasarkan channel
* [ ] Message status berubah menjadi `sent` atau `failed`
* [ ] Jika gagal, worker melakukan retry
* [ ] Retry maksimal 5 kali
* [ ] Status update dipublish ke Redis Pub/Sub

## Redis CLI Check

```bash
docker exec -it PROJECT_REDIS_CONTAINER redis-cli
```

```redis
KEYS queue:*
```

## Failure Condition

Phase gagal jika:

* Job tidak masuk queue
* Worker tidak mengambil job
* Worker crash saat provider error
* Retry tidak berjalan
* Message status tidak berubah
* Failed message tidak ditandai sebagai `failed`

---

# 11. Phase 07 — MinIO Attachment Test

## Scope

Phase ini menguji:

* Presigned upload URL
* Upload file ke MinIO
* Confirm upload
* Metadata attachment
* Signed download URL
* Preview attachment di React

## Presigned Upload Test

```bash
curl -X POST http://localhost:8080/api/attachments/presigned-upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "filename":"katalog.pdf",
    "mime_type":"application/pdf",
    "size_bytes":1200000
  }'
```

## Expected Result

* [ ] Backend mengembalikan upload URL
* [ ] Backend mengembalikan object key
* [ ] URL punya expiry
* [ ] File size divalidasi
* [ ] MIME type divalidasi

## Example Response

```json
{
  "success": true,
  "message": "Presigned upload URL created",
  "data": {
    "upload_url": "temporary-url",
    "object_key": "chat-media/2026/07/conversation-id/uuid-katalog.pdf",
    "expires_in": 900
  }
}
```

## Upload File Test

```bash
curl -X PUT "PRESIGNED_UPLOAD_URL" \
  -H "Content-Type: application/pdf" \
  --upload-file ./katalog.pdf
```

## Confirm Upload Test

```bash
curl -X POST http://localhost:8080/api/attachments/confirm-upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "message_id":"MESSAGE_ID",
    "bucket":"chat-media",
    "object_key":"chat-media/2026/07/conversation-id/uuid-katalog.pdf",
    "original_filename":"katalog.pdf",
    "mime_type":"application/pdf",
    "size_bytes":1200000
  }'
```

## Signed Download URL Test

```bash
curl http://localhost:8080/api/attachments/ATTACHMENT_ID/signed-url \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Metadata attachment tersimpan di PostgreSQL
* [ ] File tersedia di MinIO
* [ ] Signed download URL dapat dibuat
* [ ] Bucket tidak public
* [ ] File bisa dibuka melalui signed URL

## Failure Condition

Phase gagal jika:

* Presigned URL gagal dibuat
* File tidak masuk MinIO
* Metadata tidak tersimpan
* File bisa diakses tanpa signed URL
* File size invalid tetap diterima
* MIME type invalid tetap diterima

---

# 12. Phase 08 — WhatsApp Integration Test

## Scope

Phase ini menguji:

* WhatsApp provider interface
* WhatsApp webhook
* Parse payload WhatsApp
* Redis idempotency
* Simpan customer
* Simpan conversation
* Simpan message
* Send WhatsApp message

## Manual Webhook Test

Contoh payload bisa berbeda sesuai provider. Untuk WhatsApp Cloud API, gunakan payload simulasi:

```bash
curl -X POST http://localhost:8080/api/webhooks/whatsapp \
  -H "Content-Type: application/json" \
  -d '{
    "entry": [
      {
        "changes": [
          {
            "value": {
              "contacts": [
                {
                  "profile": {
                    "name": "Budi WhatsApp"
                  },
                  "wa_id": "6281234567890"
                }
              ],
              "messages": [
                {
                  "from": "6281234567890",
                  "id": "wamid.test123",
                  "timestamp": "1710000000",
                  "text": {
                    "body": "Halo, saya mau tanya katalog"
                  },
                  "type": "text"
                }
              ]
            }
          }
        ]
      }
    ]
  }'
```

## Expected Result

* [ ] Payload diterima
* [ ] Phone number berhasil diparse
* [ ] Customer dibuat atau ditemukan
* [ ] Conversation channel `whatsapp` dibuat atau ditemukan
* [ ] Message tersimpan
* [ ] Redis idempotency berjalan
* [ ] Duplicate webhook tidak membuat message dobel
* [ ] Event realtime dipublish

## Send WhatsApp Test

```bash
curl -X POST http://localhost:8080/api/conversations/CONVERSATION_ID/messages \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body":"Halo kak, berikut katalog kami.","template_id":null,"attachment_id":null}'
```

## Expected Result

* [ ] Message masuk Redis queue
* [ ] Worker memanggil WhatsApp provider
* [ ] Status menjadi `sent` jika sukses
* [ ] Status menjadi `failed` jika provider error
* [ ] Error provider tidak membuat worker crash

## Failure Condition

Phase gagal jika:

* WhatsApp payload tidak bisa diparse
* Provider message ID tidak tersimpan
* Duplicate webhook membuat pesan dobel
* Provider interface tidak modular
* Send message tidak lewat queue
* Token WhatsApp hardcode

---

# 13. Phase 09 — Template Message Test

## Scope

Phase ini menguji:

* CRUD template pesan
* Template channel
* Template picker di React
* Template digunakan saat reply

## Create Template

```bash
curl -X POST http://localhost:8080/api/templates \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Info Harga",
    "content":"Halo kak, berikut informasi harga produk kami.",
    "channel":"all",
    "attachment_id":null,
    "is_active":true
  }'
```

## Expected Result

* [ ] Template berhasil dibuat
* [ ] Template tersimpan di PostgreSQL
* [ ] Channel valid
* [ ] Template aktif

## Get Templates

```bash
curl http://localhost:8080/api/templates \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Update Template

```bash
curl -X PUT http://localhost:8080/api/templates/TEMPLATE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title":"Info Harga Updated",
    "content":"Halo kak, ini daftar harga terbaru kami.",
    "channel":"all",
    "attachment_id":null,
    "is_active":true
  }'
```

## Delete Template

```bash
curl -X DELETE http://localhost:8080/api/templates/TEMPLATE_ID \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Frontend Test

Checklist:

* [ ] Halaman template tampil
* [ ] Admin bisa membuat template
* [ ] Admin bisa edit template
* [ ] Admin bisa hapus template
* [ ] Template picker tampil di reply box
* [ ] Saat template dipilih, isi reply box berubah

## Failure Condition

Phase gagal jika:

* Template tidak tersimpan
* Template channel invalid diterima
* Template picker tidak mengisi reply box
* Endpoint bisa diakses tanpa auth

---

# 14. Phase 10 — Auto-Reply Keyword Test

## Scope

Phase ini menguji:

* CRUD auto-reply rule
* Keyword matching
* Auto-reply by text
* Auto-reply by template
* Auto-send through Redis queue

## Create Auto-Reply Rule

```bash
curl -X POST http://localhost:8080/api/auto-replies \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "keyword":"harga",
    "match_type":"contains",
    "template_id":null,
    "response_text":"Halo kak, berikut informasi harga produk kami.",
    "channel":"all",
    "is_active":true
  }'
```

## Expected Result

* [ ] Rule berhasil dibuat
* [ ] Keyword tersimpan
* [ ] Match type valid
* [ ] Channel valid
* [ ] Rule aktif

## Trigger Auto-Reply Test

Kirim pesan customer via webhook:

```bash
curl -X POST http://localhost:8080/api/webhooks/telegram \
  -H "Content-Type: application/json" \
  -d '{
    "update_id": 100002,
    "message": {
      "message_id": 200002,
      "from": {
        "id": 123456789,
        "is_bot": false,
        "first_name": "Budi",
        "username": "budi_test"
      },
      "chat": {
        "id": 123456789,
        "first_name": "Budi",
        "username": "budi_test",
        "type": "private"
      },
      "date": 1710000000,
      "text": "Saya mau tanya harga"
    }
  }'
```

## Expected Result

* [ ] Pesan customer tersimpan
* [ ] Rule keyword `harga` terdeteksi
* [ ] Message auto-reply dibuat dengan sender type `system`
* [ ] Auto-reply masuk Redis queue
* [ ] Worker mengirim auto-reply
* [ ] Histori conversation menampilkan auto-reply

## Negative Test

Kirim pesan tanpa keyword:

```bash
curl -X POST http://localhost:8080/api/webhooks/telegram \
  -H "Content-Type: application/json" \
  -d '{
    "update_id": 100003,
    "message": {
      "message_id": 200003,
      "from": {
        "id": 123456789,
        "is_bot": false,
        "first_name": "Budi",
        "username": "budi_test"
      },
      "chat": {
        "id": 123456789,
        "first_name": "Budi",
        "username": "budi_test",
        "type": "private"
      },
      "date": 1710000000,
      "text": "Terima kasih"
    }
  }'
```

Expected:

* [ ] Pesan tersimpan
* [ ] Tidak ada auto-reply dibuat

## Failure Condition

Phase gagal jika:

* Keyword tidak terdeteksi
* Keyword salah mendeteksi pesan tidak relevan
* Auto-reply tidak masuk queue
* Auto-reply tetap berjalan untuk conversation `resolved`
* Auto-reply membuat infinite loop

---

# 15. Phase 11 — Realtime Inbox Test

## Scope

Phase ini menguji:

* Redis Pub/Sub
* SSE endpoint
* React realtime listener
* Inbox update otomatis
* Message status update otomatis

## SSE Endpoint Test

```bash
curl http://localhost:8080/api/realtime/events \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

* [ ] Koneksi SSE terbuka
* [ ] Tidak langsung tertutup
* [ ] Event dapat diterima saat ada pesan baru

## New Message Event Test

1. Buka dashboard React.
2. Buka halaman inbox.
3. Kirim webhook Telegram/WhatsApp manual.
4. Lihat inbox.

Expected:

* [ ] Conversation list refresh otomatis
* [ ] Pesan baru muncul tanpa reload browser
* [ ] Toast notification muncul
* [ ] Badge unread berubah

## Message Status Event Test

1. Admin kirim reply.
2. Worker memproses queue.
3. Status message berubah.

Expected:

* [ ] UI message status berubah dari `pending` ke `sent`
* [ ] Tidak perlu refresh manual
* [ ] Jika gagal, status berubah menjadi `failed`

## Failure Condition

Phase gagal jika:

* SSE tidak connect
* Event tidak diterima frontend
* Inbox tidak update
* Browser harus reload untuk melihat pesan baru
* SSE endpoint bisa diakses tanpa auth

---

# 16. Phase 12 — Dashboard Summary Test

## Scope

Phase ini menguji:

* Dashboard summary API
* Redis cache
* Summary card di React
* Statistik pesan

## Test Command

```bash
curl http://localhost:8080/api/dashboard/summary \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Expected Result

Response berisi:

* [ ] Total pesan hari ini
* [ ] Total WhatsApp hari ini
* [ ] Total Telegram hari ini
* [ ] Unreplied conversation
* [ ] Resolved conversation
* [ ] Auto-reply sent today
* [ ] Attachment uploaded today

## Example Response

```json
{
  "success": true,
  "message": "Dashboard summary fetched",
  "data": {
    "total_today": 35,
    "whatsapp_today": 24,
    "telegram_today": 11,
    "unreplied": 7,
    "resolved": 15,
    "auto_reply_sent": 9,
    "attachments_today": 4
  }
}
```

## Redis Cache Test

1. Request endpoint pertama kali.
2. Request endpoint kedua kali.
3. Cek log backend atau Redis key.

Redis key:

```text
cache:dashboard:summary:YYYY-MM-DD
```

Expected:

* [ ] Request pertama mengambil data dari PostgreSQL
* [ ] Request kedua mengambil data dari Redis cache
* [ ] TTL cache berjalan
* [ ] Data tidak error saat cache kosong

## Frontend Test

Checklist:

* [ ] Dashboard page tampil
* [ ] Summary card tampil
* [ ] Loading state tampil
* [ ] Error state tampil jika API gagal
* [ ] Data sesuai response backend

## Failure Condition

Phase gagal jika:

* Summary API error
* Statistik salah
* Redis cache tidak bekerja
* Frontend gagal menampilkan data
* Endpoint bisa diakses tanpa auth

---

# 17. Phase 13 — Kubernetes Deployment Test

## Scope

Phase ini menguji deployment Kubernetes basic:

* Namespace
* ConfigMap
* Secret
* Deployment backend
* Deployment worker
* Deployment frontend
* Service backend
* Service frontend
* Ingress
* PostgreSQL PVC
* Redis PVC
* MinIO PVC

## Test Commands

```bash
kubectl apply -f k8s/namespace.yaml
```

```bash
kubectl apply -f k8s/
```

```bash
kubectl get pods -n cs-dashboard
```

```bash
kubectl get svc -n cs-dashboard
```

```bash
kubectl get ingress -n cs-dashboard
```

## Expected Result

* [ ] Namespace berhasil dibuat
* [ ] ConfigMap berhasil dibuat
* [ ] Secret berhasil dibuat
* [ ] Backend pod running
* [ ] Worker pod running
* [ ] Frontend pod running
* [ ] PostgreSQL pod running
* [ ] Redis pod running
* [ ] MinIO pod running
* [ ] Service backend aktif
* [ ] Service frontend aktif
* [ ] Ingress aktif
* [ ] PVC bound

## Backend Health Check in Kubernetes

```bash
kubectl port-forward svc/backend-service 8080:8080 -n cs-dashboard
```

```bash
curl http://localhost:8080/health
```

Expected:

* [ ] Health check OK
* [ ] Backend connect ke PostgreSQL
* [ ] Backend connect ke Redis
* [ ] Backend connect ke MinIO

## Failure Condition

Phase gagal jika:

* Pod crash loop
* PVC tidak bound
* Backend tidak connect ke dependency
* Ingress tidak bisa route frontend/backend
* Secret tidak terbaca
* ConfigMap tidak terbaca

---

# 18. Phase 14 — Final Review Test

## Scope

Phase ini menguji keseluruhan sistem sebelum dianggap MVP selesai.

## Full System Test

Jalankan:

```bash
docker compose down
docker compose up -d --build
```

Checklist:

* [ ] Semua container running
* [ ] Backend health OK
* [ ] Frontend bisa dibuka
* [ ] Login berhasil
* [ ] Inbox tampil
* [ ] Telegram webhook masuk
* [ ] WhatsApp webhook masuk
* [ ] Admin bisa reply
* [ ] Redis queue berjalan
* [ ] Worker berjalan
* [ ] Auto-reply berjalan
* [ ] Template bisa digunakan
* [ ] Attachment bisa upload ke MinIO
* [ ] Attachment bisa download dengan signed URL
* [ ] Realtime update berjalan
* [ ] Dashboard summary tampil
* [ ] Tidak ada error fatal di log

## Backend Test

```bash
cd backend
go test ./...
```

Expected:

* [ ] Semua test backend lulus
* [ ] Tidak ada panic
* [ ] Tidak ada race condition sederhana

## Frontend Test

```bash
cd frontend
npm run build
```

Expected:

* [ ] Build berhasil
* [ ] Tidak ada TypeScript error
* [ ] Tidak ada import error

## Security Checklist

* [ ] Password di-hash bcrypt
* [ ] JWT secret dari environment
* [ ] `.env` tidak dicommit
* [ ] Protected route memakai middleware
* [ ] Webhook memakai verify token jika tersedia
* [ ] File download memakai signed URL
* [ ] Bucket MinIO tidak public
* [ ] Login punya rate limit
* [ ] Webhook punya idempotency
* [ ] Token tidak muncul di log

## Performance Checklist

* [ ] Dashboard summary memakai Redis cache
* [ ] Message sending memakai Redis queue
* [ ] Webhook tidak menjalankan proses berat terlalu lama
* [ ] Attachment disimpan ke MinIO
* [ ] Query conversation memakai index
* [ ] Query message memakai index

## Failure Condition

Final review gagal jika:

* Core feature tidak berjalan
* Login gagal
* Inbox tidak tampil
* Webhook tidak menyimpan pesan
* Reply tidak berjalan
* Queue tidak berjalan
* Attachment tidak bisa upload/download
* Ada secret hardcode
* Build frontend gagal
* Backend crash

---

# 19. Manual End-to-End Test Scenario

## Scenario 1 — Telegram Customer Message

Steps:

1. Customer mengirim pesan ke Telegram Bot.
2. Telegram webhook masuk ke backend.
3. Backend menyimpan customer.
4. Backend membuat conversation.
5. Backend menyimpan message.
6. Dashboard menampilkan pesan baru.
7. Admin membuka conversation.
8. Admin membalas pesan.
9. Message masuk Redis queue.
10. Worker mengirim pesan ke Telegram.
11. Status message berubah menjadi `sent`.

Expected:

* [ ] Customer tampil
* [ ] Conversation tampil
* [ ] Message tampil
* [ ] Reply terkirim
* [ ] Status terkirim
* [ ] Tidak ada duplikasi message

---

## Scenario 2 — WhatsApp Customer Message

Steps:

1. Customer mengirim pesan WhatsApp.
2. WhatsApp webhook masuk ke backend.
3. Backend menyimpan customer berdasarkan phone number.
4. Backend membuat conversation channel WhatsApp.
5. Message tampil di inbox.
6. Admin membalas.
7. Worker mengirim pesan melalui WhatsApp provider.

Expected:

* [ ] Phone number tersimpan
* [ ] Conversation channel WhatsApp tampil
* [ ] Admin reply tersimpan
* [ ] Message masuk queue
* [ ] Status update berjalan

---

## Scenario 3 — Auto-Reply Keyword

Steps:

1. Admin membuat rule keyword `harga`.
2. Customer mengirim pesan “Saya mau tanya harga”.
3. Backend mendeteksi keyword.
4. Backend membuat auto-reply.
5. Auto-reply masuk queue.
6. Worker mengirim auto-reply.
7. Histori chat menampilkan auto-reply.

Expected:

* [ ] Rule terdeteksi
* [ ] Auto-reply dibuat
* [ ] Auto-reply terkirim
* [ ] Tidak infinite loop

---

## Scenario 4 — Attachment Upload

Steps:

1. Admin membuka conversation.
2. Admin memilih file PDF.
3. Frontend meminta presigned URL.
4. Frontend upload file ke MinIO.
5. Frontend confirm upload ke backend.
6. Metadata tersimpan.
7. Admin mengirim message dengan attachment.
8. Customer menerima attachment jika provider mendukung.

Expected:

* [ ] File masuk MinIO
* [ ] Metadata tersimpan
* [ ] Signed URL bisa dibuat
* [ ] File bisa dibuka
* [ ] File tidak public

---

# 20. Bug Reporting Format

Jika ada error, catat di `ERROR_LOG.md` dengan format:

```md
## ERR-001

### Date
YYYY-MM-DD

### Phase
Phase name

### Error
Describe the error.

### Steps to Reproduce
1. Step one
2. Step two
3. Step three

### Expected Result
Describe expected result.

### Actual Result
Describe actual result.

### Root Cause
Describe root cause if known.

### Fix
Describe fix.

### Status
OPEN / FIXED / NEEDS REVIEW
```

---

# 21. Final Acceptance Criteria

MVP dianggap selesai jika:

* [ ] Admin bisa login
* [ ] Admin bisa logout
* [ ] Dashboard bisa dibuka
* [ ] Telegram webhook berjalan
* [ ] WhatsApp webhook berjalan
* [ ] Inbox menampilkan pesan Telegram dan WhatsApp
* [ ] Detail conversation menampilkan histori
* [ ] Admin bisa membalas pesan
* [ ] Pesan keluar diproses via Redis queue
* [ ] Retry berjalan untuk pesan gagal
* [ ] Webhook idempotency mencegah duplikasi
* [ ] Auto-reply keyword berjalan
* [ ] Template pesan bisa dibuat dan digunakan
* [ ] Attachment tersimpan di MinIO
* [ ] Signed URL berjalan untuk attachment
* [ ] Realtime inbox notification berjalan
* [ ] Dashboard summary tampil
* [ ] Docker Compose berjalan
* [ ] Kubernetes manifest tersedia
* [ ] Tidak ada secret hardcode
* [ ] Frontend build berhasil
* [ ] Backend test berhasil
