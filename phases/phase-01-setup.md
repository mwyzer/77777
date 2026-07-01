# Phase 01 — Project Setup

## Goal

Membuat setup awal project Customer Communication Dashboard.

## Tech

- Golang
- Gin
- PostgreSQL
- Redis
- MinIO
- Docker Compose

## Tasks

1. Buat backend Golang.
2. Buat endpoint GET /health.
3. Setup config dari .env.
4. Setup koneksi PostgreSQL.
5. Setup koneksi Redis.
6. Setup koneksi MinIO.
7. Buat Dockerfile backend.
8. Buat docker-compose.yml.
9. Pastikan semua service bisa jalan.

## Acceptance Criteria

Phase ini dianggap selesai jika:

1. `docker compose up -d` berhasil.
2. Backend berjalan di port 8080.
3. Endpoint `GET /health` mengembalikan status OK.
4. PostgreSQL terkoneksi.
5. Redis terkoneksi.
6. MinIO terkoneksi.
7. Tidak ada error fatal di log backend.

## Test Command

```bash
docker compose up -d
curl http://localhost:8080/health
docker compose logs backend