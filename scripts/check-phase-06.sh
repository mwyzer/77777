#!/bin/bash
# Phase 06 — Redis Queue Worker Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-06.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="06"

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
    echo "  Acceptance criteria: Queue send message, retry, status update"
    exit 0
fi

BASE_URL="${BASE_URL:-http://localhost:8080}"
echo "[1/3] Checking Redis connectivity..."
docker compose exec -T redis redis-cli ping 2>/dev/null | grep -q PONG && echo "  PASS: Redis reachable" || echo "  WARN: Redis check failed"

echo "[2/3] Checking worker binary..."
[ -f "$ROOT_DIR/backend/cmd/worker/main.go" ] && echo "  PASS: worker/main.go exists" || echo "  WARN: worker binary not found"

echo "[3/3] Checking queue keys in Redis..."
docker compose exec -T redis redis-cli KEYS "queue:*" 2>/dev/null | head -5
echo "  INFO: Queue keys listed above"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
