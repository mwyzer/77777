#!/bin/bash
# Phase 09 — Template Message Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-09.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="09"

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
    echo "  Acceptance criteria: CRUD template and template picker"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
get_token() {
    curl -s -X POST "$BASE_URL/api/auth/login" -H "Content-Type: application/json" \
      -d '{"email":"admin@example.com","password":"admin123"}' | grep -o '"token":"[^"]*"' | head -1 | sed 's/"token":"//;s/"//'
}

TOKEN=$(get_token)
echo "[1/3] Checking template endpoints..."
TEMPLATES=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/templates" -H "Authorization: Bearer $TOKEN" 2>/dev/null)
echo "  INFO: GET /api/templates → HTTP $TEMPLATES"

echo "[2/3] Checking message_templates table..."
docker compose exec -T postgres psql -U dashboard -d dashboard -c "SELECT COUNT(*) FROM message_templates;" 2>/dev/null && echo "  PASS: table exists" || echo "  WARN: table check failed"

echo "[3/3] Checking template module..."
[ -d "$ROOT_DIR/backend/internal/modules/template" ] && echo "  PASS: template module exists" || echo "  WARN: template module not found"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
