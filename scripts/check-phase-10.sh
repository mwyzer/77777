#!/bin/bash
# Phase 10 — Auto-Reply Keyword Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-10.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="10"

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
    echo "  Acceptance criteria: Keyword rule, auto-send reply through queue"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
get_token() {
    curl -s -X POST "$BASE_URL/api/auth/login" -H "Content-Type: application/json" \
      -d '{"email":"admin@example.com","password":"admin123"}' | grep -o '"token":"[^"]*"' | head -1 | sed 's/"token":"//;s/"//'
}

TOKEN=$(get_token)
echo "[1/3] Checking auto-reply endpoints..."
RULES=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/auto-reply-rules" -H "Authorization: Bearer $TOKEN" 2>/dev/null)
echo "  INFO: GET /api/auto-reply-rules → HTTP $RULES"

echo "[2/3] Checking auto_reply_rules table..."
docker compose exec -T postgres psql -U dashboard -d dashboard -c "SELECT COUNT(*) FROM auto_reply_rules;" 2>/dev/null && echo "  PASS: table exists" || echo "  WARN: table check failed"

echo "[3/3] Checking auto-reply module..."
[ -d "$ROOT_DIR/backend/internal/modules/autoreply" ] && echo "  PASS: autoreply module exists" || echo "  WARN: autoreply module not found"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
