#!/bin/bash
# Phase 11 — Realtime Inbox Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-11.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="11"

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
    echo "  Acceptance criteria: Redis Pub/Sub and SSE/WebSocket"
    exit 0
fi

echo "[1/3] Checking Redis Pub/Sub channels..."
docker compose exec -T redis redis-cli PUBSUB CHANNELS 2>/dev/null | grep -q "pubsub:" && echo "  PASS: Redis Pub/Sub active" || echo "  WARN: No pubsub channels found"

echo "[2/3] Checking realtime module..."
[ -d "$ROOT_DIR/backend/internal/realtime" ] && echo "  PASS: realtime/ directory exists" || echo "  WARN: realtime module not found"

echo "[3/3] Checking SSE/WebSocket endpoint..."
BASE_URL="${BASE_URL:-http://localhost:8080}"
curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/realtime/events" 2>/dev/null
echo "  INFO: GET /api/realtime/events"
echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
