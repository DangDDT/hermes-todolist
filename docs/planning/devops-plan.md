# DevOps Plan

## 1. MVP Deployment Model
- single VPS
- Docker Compose
- reverse proxy via Nginx
- backend + frontend + db + proxy services

## 2. Environments
| Environment | Purpose |
|---|---|
| Local Dev | active coding and verification |
| CI | lint, test, build validation |
| Staging (recommended) | pre-prod smoke validation |
| Production | user-facing deployment |

## 3. CI/CD Baseline
### CI
- install dependencies
- build backend
- build frontend
- run tests when available
- fail fast on compile issues

### CD
- manual trigger initially
- deploy through Docker update on VPS
- smoke check after deployment

## 4. Secrets & Config
- env file template in repo
- production secrets outside repo
- JWT secret rotation procedure required
- DB credentials per environment

## 5. Observability
Minimum baseline:
- structured logs
- health endpoint
- container restart policy
- basic alert signal

Near-term recommended:
- Sentry
- uptime monitor
- Grafana/Prometheus or equivalent
- backup verification

## 6. Backup & Recovery
- scheduled PostgreSQL backups
- retention policy
- rollback image/version tagging
- deploy rollback runbook

## 7. Release Checklist
- CI green
- smoke tests green
- env validated
- migration reviewed
- rollback path ready
- issue/milestone updated on GitHub
