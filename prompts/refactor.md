You are refactoring the project safely.

Read:

PROJECT_RULES.md
PHASE_STATUS.md
docs/ARCHITECTURE.md
docs/API_CONTRACT.md
docs/TEST_PLAN.md

Task:

Refactor only the requested area.

Rules:

1. Do not change business behavior.
2. Do not add new features.
3. Do not remove working features.
4. Keep API contract compatible.
5. Keep folder structure modular.
6. Improve readability and maintainability.
7. Run or explain relevant tests.
8. Write a refactor report.

Refactor report format:

Refactor scope:
Reason:
Files changed:
Behavior changed:
Test command:
Test result:
Risk:
Rollback suggestion:

Preferred refactor principles:

- Handler should not contain business logic.
- Service should contain business logic.
- Repository should contain database query only.
- Provider should contain external API logic.
- Redis queue should be isolated.
- MinIO storage logic should be isolated.
- Frontend API calls should go through services/api.ts.