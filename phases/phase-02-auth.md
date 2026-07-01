
---

# Contoh Prompt Otomatis ke DeepSeek

Pakai prompt ini setelah kamu punya folder `phases`.

```text
Kamu adalah AI coding agent untuk project Customer Communication Dashboard.

Baca file berikut:
1. docs/BRD.md
2. docs/SRS.md
3. docs/ARCHITECTURE.md
4. phases/phase-01-setup.md

Kerjakan hanya Phase 01.

Aturan:
1. Jangan lanjut ke Phase 02 sebelum Phase 01 selesai.
2. Ikuti acceptance criteria.
3. Setelah coding selesai, jalankan atau jelaskan test command.
4. Jika ada error, perbaiki sampai lolos.
5. Setelah Phase 01 lolos, tulis laporan:
   - File yang dibuat
   - File yang diubah
   - Cara menjalankan
   - Hasil test
   - Status: PASSED atau FAILED
6. Jika PASSED, baru boleh rekomendasikan lanjut Phase 02.

Jangan membuat fitur auth, inbox, Telegram, WhatsApp, Redis queue, MinIO attachment, atau Kubernetes dulu kecuali yang diminta di Phase 01.