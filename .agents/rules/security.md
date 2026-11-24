# Security

Practice scenarios use synthetic logs, account identifiers, hostnames and metrics. Review pasted evidence for credentials, session tokens and customer data before it enters fixtures, snapshots, prompts or documentation. The ignored environment-file patterns in `.gitignore` cover specific filenames. Inspect the actual diff for accidental disclosures.

Treat learner answers and imported evidence as untrusted text. Render plain text with React's normal escaping. Keep AWS CLI and kubectl snippets in exercises inert. A learner selecting a diagnostic step should reveal simulated evidence.

If a future feature introduces private progress or credentials, identify which server code enforces ownership and which fields cross into the browser. Use the [security audit workflow](../skills/security-audit/SKILL.md) for reviews.
