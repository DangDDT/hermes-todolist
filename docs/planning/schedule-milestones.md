# Schedule & Milestones

## 1. Planning Assumptions
- Baseline team: 3–5 contributors equivalent
- Sprint cadence: 1–2 weeks
- MVP target: phased delivery, not big bang

---

## 2. Milestones
| Milestone | Target | Exit Criteria |
|---|---|---|
| M1 Planning Baseline Approved | 2026-07-01 | Planning pack approved by PO |
| M2 Design Baseline Approved | 2026-07-03 | SDD + diagrams aligned |
| M3 Backend MVP Complete | 2026-07-08 | Auth + task CRUD + build pass |
| M4 Frontend MVP Complete | 2026-07-10 | Auth + task flows UI complete |
| M5 Test Baseline Complete | 2026-07-13 | Unit/integration/component suites available |
| M6 Deployable Release Candidate | 2026-07-15 | CI green + Docker deploy baseline |
| M7 Monitoring Baseline | 2026-07-16 | Logs + health + alerting minimum viable |

---

## 3. Indicative Schedule by Workstream
### Week 1
- Planning pack approval
- Architecture/design baseline
- Backend core and frontend shell
- Database and repositories

### Week 2
- Complete frontend flows
- Testing implementation
- Swagger and docs alignment
- Monitoring and release hardening

---

## 4. Execution Mapping
| Workstream | Start | End | Depends On |
|---|---|---|---|
| Planning | 2026-06-30 | 2026-07-01 | none |
| Design | 2026-06-30 | 2026-07-03 | planning inputs |
| Backend | 2026-07-01 | 2026-07-08 | design baseline |
| Frontend | 2026-07-01 | 2026-07-10 | API/design baseline |
| Testing | 2026-07-08 | 2026-07-13 | implementation baseline |
| Deploy/Hardening | 2026-07-12 | 2026-07-16 | CI + test baseline |

---

## 5. GitHub Mapping
- Milestone: MVP v1.0 = umbrella delivery milestone
- Project Board columns = Todo / In Progress / Review / Done
- Burndown tracked via wiki + issue state
