# Budget & Resource Plan

## 1. Team Assumption
Planning baseline sử dụng giả định 3–5 contributor equivalents:
- 1 Product Owner / Project Coordinator
- 1 Backend Engineer
- 1 Frontend Engineer
- 1 QA / shared testing role
- 1 DevOps / shared platform role

Trong thực tế, một người có thể kiêm nhiều vai trò.

---

## 2. Role Allocation
| Role | Primary Responsibilities |
|---|---|
| PO | scope, priority, acceptance, business decisions |
| Backend | auth, API, repository, DB integration |
| Frontend | UI, forms, query state, UX states |
| QA | test cases, regression, bug validation |
| DevOps | CI/CD, Docker, monitoring, deployment |

---

## 3. Effort Estimate
| Workstream | Est. Person-Days |
|---|---|
| Planning & documentation | 3 |
| Architecture & design | 3 |
| Backend MVP | 5 |
| Frontend MVP | 5 |
| Testing baseline | 3 |
| DevOps & deployment baseline | 3 |
| Buffer / rework | 3 |
| **Total** | **25 person-days** |

---

## 4. Infrastructure Budget (MVP)
| Item | Monthly Estimate |
|---|---|
| VPS | $10–20 |
| Domain | $1–2 (amortized) |
| DB backups / object storage | $0–5 |
| Monitoring / error tracking | $0–10 |
| CI/CD | $0–10 |
| **Total Monthly MVP** | **~$11–47** |

---

## 5. Tooling Budget Notes
- GitHub: existing baseline sufficient for MVP.
- Docker / PostgreSQL / Nginx: open-source.
- Optional paid services should be justified only if they reduce delivery risk materially.

---

## 6. Resource Risks
- One person wearing too many hats may create schedule compression risk.
- QA may become bottleneck late if tests are delayed.
- DevOps hardening often slips if not planned explicitly.
