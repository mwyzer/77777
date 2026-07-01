#!/bin/bash
# Phase 05 — Telegram Integration Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-05.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="05"

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
    echo "  Acceptance criteria: Telegram webhook, receive message, reply message"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
echo "[1/4] Checking backend health..."
curl -sf "$BASE_URL/health" > /dev/null && echo "  PASS: Backend running" || echo "  FAIL: Backend not reachable"

echo "[2/4] Checking webhook endpoint exists..."
WEBHOOK=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/webhooks/telegram" -X POST -H "Content-Type: application/json" -d '{}')
echo "  INFO: POST /api/webhooks/telegram → HTTP $WEBHOOK"

echo "[3/4] Checking TELEGRAM_BOT_TOKEN env..."
if [ -n "$TELEGRAM_BOT_TOKEN" ]; then
    echo "  PASS: TELEGRAM_BOT_TOKEN is set"
else
    echo "  WARN: TELEGRAM_BOT_TOKEN not set (check .env)"
fi

echo "[4/4] Checking Telegram provider module..."
[ -d "$ROOT_DIR/backend/internal/providers" ] && echo "  PASS: providers/ directory exists" || echo "  WARN: providers/ not found"
echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
