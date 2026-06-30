# RACI Matrix

## 1. Purpose
RACI giúp PO nhìn rõ ai:
- Responsible (R): người thực thi chính
- Accountable (A): người chịu trách nhiệm cuối cùng
- Consulted (C): người cần được hỏi ý kiến
- Informed (I): người cần được cập nhật

---

## 2. Roles
| Role | Description |
|---|---|
| PO | Product Owner / final business decision maker |
| Tada | Delivery coordinator / documentation / execution support |
| Backend Engineer | API, domain, DB integration |
| Frontend Engineer | web UI, forms, UX states |
| QA | testing, regression, release validation |
| DevOps | CI/CD, infra, deploy, monitoring |

---

## 3. Delivery RACI
| Work Item | PO | Tada | Backend | Frontend | QA | DevOps |
|---|---|---|---|---|---|---|
| Product vision & priorities | A | C | I | I | I | I |
| PRD / scope approval | A | R | I | I | I | I |
| Planning pack drafting | C | A/R | C | C | C | C |
| Architecture baseline | C | R | A/R | C | I | C |
| Database design | C | R | A/R | I | I | C |
| API design | C | R | A/R | C | I | I |
| Backend implementation | I | C | A/R | I | I | I |
| Frontend implementation | I | C | C | A/R | I | I |
| Test case design | C | R | C | C | A/R | I |
| Automated tests implementation | I | C | R | R | A | I |
| CI/CD pipeline | I | C | C | C | I | A/R |
| Deployment approval | A | C | I | I | C | R |
| Risk register maintenance | C | A/R | C | C | C | C |
| Backlog grooming | A | R | C | C | C | C |
| Issue/status logging on GitHub | I | A/R | C | C | I | I |
| Release readiness decision | A | R | C | C | C | C |

---

## 4. Agile Ceremonies RACI
| Ceremony / Activity | PO | Tada | Backend | Frontend | QA | DevOps |
|---|---|---|---|---|---|---|
| Sprint planning | A | R | C | C | C | C |
| Daily status update | I | A/R | R | R | R | R |
| Scope change review | A | R | C | C | C | C |
| Sprint review | A | R | C | C | C | C |
| Retro / lessons learned | C | R | C | C | C | C |

---

## 5. Notes
- Trong giai đoạn hiện tại, một người có thể kiêm nhiều vai trò.
- Dù implementation có thể do Tada/subagents hỗ trợ, quyền chốt business priority vẫn thuộc PO.
- Với mọi thay đổi ảnh hưởng scope, acceptance hoặc risk profile, PO phải được informed tối thiểu, và thường là Accountable.
