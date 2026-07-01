#!/bin/bash
# Phase 12 — Dashboard Summary Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-12.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="12"

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
    echo "  Acceptance criteria: Message statistics and Redis cache"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
get_token() {
    curl -s -X POST "$BASE_URL/api/auth/login" -H "Content-Type: application/json" \
      -d '{"email":"admin@example.com","password":"admin123"}' | grep -o '"token":"[^"]*"' | head -1 | sed 's/"token":"//;s/"//'
}

TOKEN=$(get_token)
echo "[1/3] Checking dashboard summary endpoint..."
SUMMARY=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/dashboard/summary" -H "Authorization: Bearer $TOKEN" 2>/dev/null)
echo "  INFO: GET /api/dashboard/summary → HTTP $SUMMARY"

echo "[2/3] Checking Redis cache keys..."
docker compose exec -T redis redis-cli KEYS "cache:dashboard:*" 2>/dev/null | head -5
echo "  INFO: Dashboard cache keys listed above"

echo "[3/3] Checking dashboard module..."
[ -d "$ROOT_DIR/backend/internal/modules/dashboard" ] && echo "  PASS: dashboard module exists" || echo "  WARN: dashboard module not found"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
