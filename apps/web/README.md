# Incident practice

From the repository root, start the website and Go API together. Press Ctrl+C to stop both.

```sh
task dev
```

Open http://localhost:3000 and select an incident. Expand the evidence, draft a response and reveal the worked diagnosis. Notes last for the current page visit.

The web server reads the v1 API at http://127.0.0.1:8080 by default. Set `INCIDENT_API_ORIGIN` on the web server to use another API host.
