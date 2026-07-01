Fix Error Prompt

You are fixing an error in the current phase.

Read these files:

PROJECT_RULES.md
PHASE_STATUS.md
ERROR_LOG.md
docs/TEST_PLAN.md
Current phase file in /phases

Task:

Fix the current error only.

Rules:

1. Do not add new features.
2. Do not skip the current phase.
3. Do not rewrite the whole project unless necessary.
4. Identify the root cause.
5. Apply the smallest safe fix.
6. Run or explain the test command again.
7. Update ERROR_LOG.md.
8. Update PHASE_STATUS.md if the phase status changes.
9. Write a fix report.

Fix report format:

Error:
Root cause:
Files changed:
Fix applied:
Test command:
Test result:
Status:

If the error cannot be fixed, explain:

1. What was attempted
2. What still fails
3. What information is needed