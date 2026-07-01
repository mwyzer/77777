#!/bin/bash
# Phase 13 — Kubernetes Deployment Check Script
# Reads PHASE_STATUS.md for phase info automatically
# Usage: bash scripts/check-phase-13.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STATUS_FILE="$ROOT_DIR/project/PHASE_STATUS.md"
PHASE="13"

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
    echo "  Acceptance criteria: Manifest basic for backend, frontend, worker, Redis, MinIO, PostgreSQL"
    exit 0
fi

echo "[1/4] Checking kubectl availability..."
kubectl version --client 2>/dev/null | head -1 && echo "  PASS: kubectl available" || echo "  WARN: kubectl not found"

echo "[2/4] Checking k8s manifest directory..."
[ -d "$ROOT_DIR/k8s" ] && echo "  PASS: k8s/ directory exists" || echo "  WARN: k8s/ directory not found"

echo "[3/4] Checking manifest files..."
for kind in namespace configmap secret deployment service ingress pvc; do
    count=$(find "$ROOT_DIR/k8s" -name "*.$kind.yaml" 2>/dev/null | wc -l)
    echo "  INFO: $kind manifests: $count file(s)"
done

echo "[4/4] Checking helm chart..."
[ -d "$ROOT_DIR/k8s/helm" ] && echo "  PASS: helm chart exists" || echo "  INFO: No helm chart (optional)"

echo ""
echo "========================================"
echo "Phase $PHASE Check Complete"
echo "========================================"
