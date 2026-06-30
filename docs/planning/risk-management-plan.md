# Risk Management Plan

## 1. Risk Method
- Probability: Low / Medium / High
- Impact: Low / Medium / High
- Score priority based on combined severity
- Review cadence: mỗi sprint và khi có scope change lớn

---

## 2. Risk Register
| ID | Risk | Category | Probability | Impact | Mitigation | Contingency |
|---|---|---|---|---|---|---|
| R-01 | Scope creep từ roadmap features tràn vào MVP | Product/Delivery | High | High | Freeze MVP scope, route extras to backlog | PO re-baseline scope |
| R-02 | Frontend và backend contract lệch nhau | Technical | Medium | High | API spec + issue traceability + smoke tests | Patch contracts before feature close |
| R-03 | Thiếu test coverage làm regression tăng | Quality | High | High | Define test plan early, add tests before release gate | Freeze new features, focus stabilization |
| R-04 | Deployment baseline không đủ observability | Ops | Medium | High | Add monitoring/logging milestone before release | Manual log review + hotfix runbook |
| R-05 | Auth/security flaws trong JWT/session | Security | Medium | High | Review cookie flags, validation, expiry, rate limiting | Rotate secret, patch auth flow |
| R-06 | Single-VPS failure impacts service availability | Infra | Medium | Medium | Backups + restart policy + health checks | Restore from backup / redeploy |
| R-07 | Documentation drift vs implementation | Process | High | Medium | Update docs with each milestone/issue closure | Reconciliation review before release |
| R-08 | Planner/implementer context loss due to many parallel tasks | Delivery | Medium | Medium | Use GitHub issue granularity + wiki logs | Mid-sprint review and regroup |

---

## 3. Top Priority Risks
1. Scope creep
2. Insufficient testing
3. Contract mismatch
4. Auth/security defects

---

## 4. Risk Response Owners
| Category | Owner |
|---|---|
| Product / Scope | PO |
| Technical / API | Engineering |
| Security | Engineering + PO review |
| Ops / Deployment | DevOps |
| Documentation drift | Tada / project maintainer |
