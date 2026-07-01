#!/bin/bash
# Phase 14 — Final Review Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-14.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="14"

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
    echo "  Acceptance criteria: Refactor, security review, performance check"
    exit 0
fi

echo "[1/5] Checking all phases are PASSED..."
TOTAL=$(grep -c "| [0-9]" "$STATUS_FILE")
PASSED=$(grep "PASSED" "$STATUS_FILE" | grep -c "| [0-9]")
echo "  INFO: $PASSED / $TOTAL phases PASSED"

echo "[2/5] Running Go vet..."
cd "$ROOT_DIR/backend" && go vet ./... 2>&1 && echo "  PASS: go vet clean" || echo "  WARN: go vet found issues"

echo "[3/5] Checking .env not committed..."
git ls-files --error-unmatch .env 2>/dev/null && echo "  WARN: .env is tracked by git!" || echo "  PASS: .env is in .gitignore"

echo "[4/5] Checking for hardcoded secrets..."
grep -r "password\s*=" "$ROOT_DIR/backend" --include="*.go" 2>/dev/null | grep -v "_test.go" | grep -v "PasswordHash" | grep -v "password_hash" | head -3
echo "  INFO: Manual review recommended for above"

echo "[5/5] Checking Docker image sizes..."
docker compose images 2>/dev/null | tail -n +2
echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
