# Phase 04 — React Dashboard Base

## Objective

Membuat React frontend dengan login UI, layout sidebar, protected route, dan halaman inbox + customers.

## Acceptance Criteria

- [x] Project Vite + React + TypeScript ter-scaffold
- [x] Tailwind CSS terkonfigurasi
- [x] Login page dengan form email/password
- [x] Auth store Zustand (token localStorage, login, logout)
- [x] API service Axios dengan JWT interceptor
- [x] Layout dengan sidebar navigasi (Inbox, Customers)
- [x] ProtectedRoute — redirect ke /login jika tidak ada token
- [x] InboxPage — list conversations dengan customer info, channel badge, status
- [x] ConversationDetailPage — chat view dengan reply box
- [x] CustomersPage — table list customers dengan provider badge
- [x] React Router routing lengkap
- [x] TanStack Query untuk server state
- [x] Dockerfile + nginx.conf untuk production build
- [x] Frontend service di docker-compose.yml

## Files Created

- `frontend/package.json`
- `frontend/vite.config.ts`
- `frontend/tsconfig.json`
- `frontend/tailwind.config.js`
- `frontend/postcss.config.js`
- `frontend/index.html`
- `frontend/Dockerfile`
- `frontend/nginx.conf`
- `frontend/src/main.tsx`
- `frontend/src/App.tsx`
- `frontend/src/index.css`
- `frontend/src/vite-env.d.ts`
- `frontend/src/types/index.ts`
- `frontend/src/services/api.ts`
- `frontend/src/stores/authStore.ts`
- `frontend/src/components/ProtectedRoute.tsx`
- `frontend/src/components/Layout.tsx`
- `frontend/src/pages/LoginPage.tsx`
- `frontend/src/pages/InboxPage.tsx`
- `frontend/src/pages/ConversationDetailPage.tsx`
- `frontend/src/pages/CustomersPage.tsx`

## Files Changed

- `docker-compose.yml` — Added frontend service (port 5173)
- `project/PHASE_STATUS.md` — Phase 04: TODO → IN_PROGRESS

## Routes

| Path                 | Component              | Auth      |
| -------------------- | ---------------------- | --------- |
| `/login`             | LoginPage              | Public    |
| `/`                  | InboxPage              | Protected |
| `/conversations/:id` | ConversationDetailPage | Protected |
| `/customers`         | CustomersPage          | Protected |
| `*`                  | Redirect to `/`        | —         |

## Test Commands

```bash
# Install deps
cd frontend && npm install

# Dev server
npm run dev

# Production build
npm run build

# Docker build
docker compose build frontend
```
