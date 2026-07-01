# git-push-phase.ps1
# Auto commit + push when a phase is PASSED.
# Usage: .\scripts\git-push-phase.ps1 -Phase "03" -Name "Core Inbox"

param(
    [Parameter(Mandatory=$true)][string]$Phase,
    [Parameter(Mandatory=$true)][string]$Name
)

Set-Location $PSScriptRoot\..
$branch = "main"

# Verify phase is PASSED
$phaseFile = "project\PHASE_STATUS.md"
$content = Get-Content $phaseFile -Raw
$pattern = "\| $Phase\s+\|.*\| PASSED \|"
if ($content -notmatch $pattern) {
    Write-Host "[SKIP] Phase $Phase is NOT marked PASSED. Nothing to push." -ForegroundColor Yellow
    exit 0
}

# Build commit message
$commitMsg = "feat(phase-$Phase): $Name"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host " Phase $Phase — $Name : PASSED" -ForegroundColor Green
Write-Host " Commit : $commitMsg" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan

# Stage all changes
git add -A
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] git add failed" -ForegroundColor Red
    exit 1
}

# Check if there is anything to commit
$status = git status --porcelain
if (-not $status) {
    Write-Host "[INFO] Nothing to commit." -ForegroundColor Yellow
    exit 0
}

# Commit
git commit -m $commitMsg
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] git commit failed" -ForegroundColor Red
    exit 1
}

# Push
git push origin $branch
if ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] git push failed" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "[OK] Phase $Phase pushed to GitHub!" -ForegroundColor Green
Write-Host ""
