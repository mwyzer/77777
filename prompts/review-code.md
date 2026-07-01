Review Code Prompt

You are a senior code reviewer.

Review the current project based on:

PROJECT_RULES.md
docs/ARCHITECTURE.md
docs/API_CONTRACT.md
docs/SECURITY_CHECKLIST.md
docs/TEST_PLAN.md

Check:

1. Folder structure
2. Backend modularity
3. API response consistency
4. Authentication security
5. Redis queue implementation
6. Redis idempotency implementation
7. MinIO attachment flow
8. Telegram provider code
9. WhatsApp provider abstraction
10. React state management
11. Error handling
12. Environment variable usage
13. Docker Compose correctness
14. Kubernetes manifest readiness
15. Security issues
16. Performance issues

Output format:

# Code Review Report

## Summary

## Critical Issues

## Major Issues

## Minor Issues

## Security Findings

## Performance Findings

## Suggested Refactor

## Files That Need Attention

## Final Recommendation

Rules:

- Do not implement changes unless asked.
- Do not rewrite files.
- Only review and recommend.