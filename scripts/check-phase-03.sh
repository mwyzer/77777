#!/bin/bash
# Phase 03 — Core Inbox Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-03.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="03"
BASE_URL="${BASE_URL:-http://localhost:8080}"

# ── Parse PHASE_STATUS.md ──
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

# ── Get auth token ──
get_token() {
    LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
      -H "Content-Type: application/json" \
      -d '{"email":"admin@example.com","password":"admin123"}')
    if echo "$LOGIN" | grep -q '"success":true'; then
        echo "$LOGIN" | grep -o '"token":"[^"]*"' | head -1 | sed 's/"token":"//;s/"//'
    else
        echo ""
    fi
}

echo "[1/7] Checking health endpoint..."
HEALTH=$(curl -s "$BASE_URL/health" 2>/dev/null)
if echo "$HEALTH" | grep -q '"success":true'; then
    echo "  PASS: Health endpoint OK"
else
    echo "  FAIL: Cannot reach API. Run: docker compose up -d"
    exit 1
fi

echo ""
echo "[2/7] Getting auth token..."
TOKEN=$(get_token)
if [ -z "$TOKEN" ]; then
    echo "  FAIL: Cannot login. Check admin credentials."
    exit 1
fi
echo "  PASS: Token obtained (${TOKEN:0:15}...)"

echo ""
echo "[3/7] GET /api/inbox/conversations..."
CONV=$(curl -s "$BASE_URL/api/inbox/conversations" -H "Authorization: Bearer $TOKEN")
if echo "$CONV" | grep -q '"success":true'; then
    TOTAL=$(echo "$CONV" | grep -o '"total":[0-9]*' | head -1 | cut -d: -f2)
    echo "  PASS: Conversations endpoint OK (total: ${TOTAL:-0})"
else
    echo "  FAIL: Conversations endpoint returned error"
    echo "  Response: ${CONV:0:200}"
fi

echo ""
echo "[4/7] GET /api/inbox/customers..."
CUST=$(curl -s "$BASE_URL/api/inbox/customers" -H "Authorization: Bearer $TOKEN")
if echo "$CUST" | grep -q '"success":true'; then
    TOTAL=$(echo "$CUST" | grep -o '"total":[0-9]*' | head -1 | cut -d: -f2)
    echo "  PASS: Customers endpoint OK (total: ${TOTAL:-0})"
else
    echo "  FAIL: Customers endpoint returned error"
    echo "  Response: ${CUST:0:200}"
fi

echo ""
echo "[5/7] Test unauthorized access (no token)..."
NO_AUTH=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/inbox/conversations")
if [ "$NO_AUTH" = "401" ]; then
    echo "  PASS: Unauthorized request rejected (401)"
else
    echo "  FAIL: Expected 401, got $NO_AUTH"
fi

echo ""
echo "[6/7] GET non-existent conversation..."
NOT_FOUND=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/inbox/conversations/00000000-0000-0000-0000-000000000000" -H "Authorization: Bearer $TOKEN")
if [ "$NOT_FOUND" = "404" ] || [ "$NOT_FOUND" = "200" ]; then
    echo "  PASS: Non-existent conversation handled (HTTP $NOT_FOUND)"
else
    echo "  WARN: Unexpected HTTP status $NOT_FOUND"
fi

echo ""
echo "[7/7] POST reply to non-existent conversation..."
REPLY=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/api/inbox/conversations/00000000-0000-0000-0000-000000000000/messages" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content":"test"}')
if [ "$REPLY" = "404" ]; then
    echo "  PASS: Reply to non-existent conversation rejected (404)"
else
    echo "  WARN: Unexpected HTTP status $REPLY"
fi

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
