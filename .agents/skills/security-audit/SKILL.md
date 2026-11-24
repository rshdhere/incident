---
name: security-audit
description: Review an incident practice change for exposed data, unsafe rendering and misplaced trust between learners, scenario content and the server.
---

# Audit a change

1. Identify the diff, the affected route or package, and which inputs a learner or scenario author can control. Read `.agents/rules/security.md`.
2. Trace each input to its use. Check how the app renders log samples, Markdown, pasted notes and links. Inspect client components and public assets for private data reaching the browser. Report credential findings by location with values redacted.
3. Where storage or server mutations exist, inspect validation, resource ownership, and which records a learner can read or modify. For infrastructure exercises, verify that actions represented as simulations only affect the exercise state.
4. Deliver findings with file locations, affected data, user impact, proposed correction and a focused defensive check. State which paths you inspected and any cases you could only assess statically.
