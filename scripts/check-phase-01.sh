#!/bin/bash
# Phase 01 — Project Setup Check Script
# Usage: bash scripts/check-phase-01.sh

echo "========================================"
echo "Phase 01 — Project Setup Check"
echo "========================================"

echo ""
echo "[1/5] Checking docker compose services..."
if docker compose ps 2>/dev/null | grep -q "Up"; then
    echo "  PASS: docker compose services running"
else
    echo "  FAIL: docker compose services not running"
    echo "  Run: docker compose up -d"
    exit 1
fi

echo ""
echo "[2/5] Checking backend container..."
if docker compose ps backend 2>/dev/null | grep -q "Up"; then
    echo "  PASS: Backend container is up"
else
    echo "  FAIL: Backend container not running"
    exit 1
fi

echo ""
echo "[3/5] Checking PostgreSQL container..."
if docker compose ps postgres 2>/dev/null | grep -q "Up"; then
    echo "  PASS: PostgreSQL container is up"
else
    echo "  FAIL: PostgreSQL container not running"
    exit 1
fi

echo ""
echo "[4/5] Checking Redis container..."
if docker compose ps redis 2>/dev/null | grep -q "Up"; then
    echo "  PASS: Redis container is up"
else
    echo "  FAIL: Redis container not running"
    exit 1
fi

echo ""
echo "[5/5] Checking MinIO container..."
if docker compose ps minio 2>/dev/null | grep -q "Up"; then
    echo "  PASS: MinIO container is up"
else
    echo "  FAIL: MinIO container not running"
    exit 1
fi

echo ""
echo "[6/6] Checking health endpoint..."
HEALTH=$(curl -s http://localhost:8080/health 2>/dev/null)
if echo "$HEALTH" | grep -q '"success":true'; then
    echo "  PASS: Health endpoint OK"
    echo "  Response: $HEALTH"
else
    echo "  FAIL: Health endpoint not OK"
    echo "  Response: $HEALTH"
    exit 1
fi

echo ""
echo "========================================"
echo "Phase 01 — ALL CHECKS PASSED!"
echo "========================================"
