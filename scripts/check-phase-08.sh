#!/bin/bash
# Phase 08 — WhatsApp Integration Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-08.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="08"

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
    echo "  Acceptance criteria: WhatsApp webhook, provider interface, send message"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
echo "[1/3] Checking webhook endpoint..."
WEBHOOK=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/webhooks/whatsapp" -X POST -H "Content-Type: application/json" -d '{}')
echo "  INFO: POST /api/webhooks/whatsapp → HTTP $WEBHOOK"

echo "[2/3] Checking WHATSAPP_ACCESS_TOKEN env..."
[ -n "$WHATSAPP_ACCESS_TOKEN" ] && echo "  PASS: WHATSAPP_ACCESS_TOKEN set" || echo "  WARN: WHATSAPP_ACCESS_TOKEN not set"

echo "[3/3] Checking provider interface..."
[ -d "$ROOT_DIR/backend/internal/providers" ] && echo "  PASS: providers/ directory exists" || echo "  WARN: providers/ not found"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
