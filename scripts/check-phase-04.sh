#!/bin/bash
# Phase 04 — React Dashboard Base Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-04.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="04"
BASE_URL="${BASE_URL:-http://localhost:8080}"
FRONTEND_URL="${FRONTEND_URL:-http://localhost:5173}"

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

echo "[1/7] Checking backend health..."
HEALTH=$(curl -s "$BASE_URL/health" 2>/dev/null)
if echo "$HEALTH" | grep -q '"success":true'; then
    echo "  PASS: Backend health OK"
else
    echo "  FAIL: Cannot reach backend API"
    exit 1
fi

echo ""
echo "[2/7] Checking frontend availability..."
FRONTEND_HTTP=$(curl -s -o /dev/null -w "%{http_code}" "$FRONTEND_URL" 2>/dev/null)
if [ "$FRONTEND_HTTP" = "200" ]; then
    echo "  PASS: Frontend serving (HTTP 200)"
else
    echo "  WARN: Frontend HTTP $FRONTEND_HTTP — is Docker running? (docker compose up -d)"
fi

echo ""
echo "[3/7] Checking frontend serves index.html..."
INDEX=$(curl -s "$FRONTEND_URL" 2>/dev/null)
if echo "$INDEX" | grep -q '<div id="root">'; then
    echo "  PASS: React root div found in HTML"
else
    echo "  WARN: React root not found — frontend may not be running"
fi

echo ""
echo "[4/7] Checking API proxy (frontend → backend)..."
PROXY=$(curl -s "$FRONTEND_URL/api/health" 2>/dev/null)
if echo "$PROXY" | grep -q '"success":true'; then
    echo "  PASS: API proxy working (frontend → backend /api/health)"
else
    echo "  WARN: API proxy may not be configured"
fi

echo ""
echo "[5/7] Checking frontend static assets..."
ASSETS=$(curl -s -o /dev/null -w "%{http_code}" "$FRONTEND_URL/assets/" 2>/dev/null)
echo "  INFO: Assets path HTTP $ASSETS"

echo ""
echo "[6/7] Checking frontend Docker container..."
if docker compose ps frontend 2>/dev/null | grep -q "Up"; then
    echo "  PASS: Frontend Docker container is up"
elif docker compose ps frontend 2>/dev/null | grep -q "Exit"; then
    echo "  FAIL: Frontend container exited"
    echo "  Check logs: docker compose logs frontend"
else
    echo "  WARN: Frontend container status unknown — is Docker running?"
fi

echo ""
echo "[7/7] Checking backend API auth endpoint..."
TOKEN=$(get_token)
if [ -n "$TOKEN" ]; then
    echo "  PASS: Auth API working (login OK)"
else
    echo "  FAIL: Cannot authenticate"
fi

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
