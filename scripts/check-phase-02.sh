#!/bin/bash
# Phase 02 — Authentication Check Script
# Usage: bash scripts/check-phase-02.sh

echo "========================================"
echo "Phase 02 — Authentication Check"
echo "========================================"

BASE_URL="http://localhost:8080"

echo ""
echo "[1/6] Check health endpoint..."
HEALTH=$(curl -s "$BASE_URL/health" 2>/dev/null)
if echo "$HEALTH" | grep -q '"success":true'; then
    echo "  PASS: Health endpoint OK"
else
    echo "  FAIL: Cannot reach API. Is docker compose running?"
    exit 1
fi

echo ""
echo "[2/6] Test login with valid credentials..."
LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password"}')
if echo "$LOGIN" | grep -q '"success":true'; then
    echo "  PASS: Login successful"
    TOKEN=$(echo "$LOGIN" | grep -o '"token":"[^"]*"' | head -1 | sed 's/"token":"//' | sed 's/"//')
    echo "  Token: ${TOKEN:0:20}..."
else
    echo "  FAIL: Login failed"
    echo "  Response: $LOGIN"
    exit 1
fi

echo ""
echo "[3/6] Test login with invalid password..."
LOGIN_FAIL=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"wrongpass"}')
if echo "$LOGIN_FAIL" | grep -q '"success":false'; then
    echo "  PASS: Invalid password rejected (401)"
else
    echo "  FAIL: Invalid password should be rejected"
fi

echo ""
echo "[4/6] Test GET /api/auth/me with valid token..."
ME=$(curl -s "$BASE_URL/api/auth/me" -H "Authorization: Bearer $TOKEN")
if echo "$ME" | grep -q '"success":true'; then
    echo "  PASS: Me endpoint OK"
    echo "  Response: $ME"
else
    echo "  FAIL: Me endpoint failed"
    echo "  Response: $ME"
fi

echo ""
echo "[5/6] Test GET /api/auth/me without token..."
ME_NO=$(curl -s "$BASE_URL/api/auth/me")
if echo "$ME_NO" | grep -q '"success":false'; then
    echo "  PASS: Request without token rejected"
else
    echo "  FAIL: Request without token should be rejected"
fi

echo ""
echo "[6/6] Test GET /api/auth/me with invalid token..."
ME_BAD=$(curl -s "$BASE_URL/api/auth/me" -H "Authorization: Bearer invalidtoken123")
if echo "$ME_BAD" | grep -q '"success":false'; then
    echo "  PASS: Invalid token rejected"
else
    echo "  FAIL: Invalid token should be rejected"
fi

echo ""
echo "========================================"
echo "Phase 02 — ALL CHECKS PASSED!"
echo "========================================"
