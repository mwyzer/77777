#!/bin/bash
# Phase 07 — MinIO Attachment Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-07.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="07"

LINE=$(grep "| $PHASE " "$STATUS_FILE" | head -1)
PHASE_NAME=$(echo "$LINE" | sed 's/| */|/g' | cut -d'|' -f3 | xargs)
PHASE_STATUS=$(echo "$LINE" | sed 's/| */|/g' | cut -d'|' -f4 | xargs)
PHASE_PUSHED=$(echo "$LINE" | sed 's/| */|/g' | cut -d'|' -f5 | xargs)

echo "========================================"
echo "Phase $PHASE — $PHASE_NAME Check"
echo "========================================"
echo "Status : $PHASE_STATUS"
echo "Pushed : $PHASE_PUSHED"
echo ""

if [ "$PHASE_STATUS" = "TODO" ]; then
    echo "[INFO] Phase $PHASE is still TODO — no implementation yet."
    echo "  Acceptance criteria: Presigned upload, signed download, attachment metadata"
    exit 0
fi

echo "[1/3] Checking MinIO connectivity..."
curl -sf http://localhost:9000/minio/health/live > /dev/null && echo "  PASS: MinIO reachable" || echo "  WARN: MinIO not reachable"

echo "[2/3] Checking MinIO bucket..."
docker compose exec -T minio mc ls local/chat-media 2>/dev/null && echo "  PASS: chat-media bucket exists" || echo "  WARN: bucket check failed"

echo "[3/3] Checking storage module..."
[ -d "$ROOT_DIR/backend/internal/storage" ] && echo "  PASS: storage/ directory exists" || echo "  WARN: storage module not found"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
