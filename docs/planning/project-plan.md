# Project Plan — Hermes TodoList

## 1. Document Control
- Project: Hermes TodoList
- Owner: DangDDT
- Product Owner: DangDDT
- Delivery Support: Tada (Hermes Agent)
- Version: 1.0
- Status: Approved for Planning Baseline
- Date: 2026-06-30

---

## 2. Project Overview
Hermes TodoList là một task management platform hướng tới team nhỏ 5–20 người, ưu tiên self-host, lightweight, nhanh, dễ dùng và có khả năng mở rộng dần lên mobile, desktop, và SaaS roadmap gần.

Mục tiêu của dự án không chỉ là làm ra một MVP chạy được, mà còn tạo nền tảng kỹ thuật và quy trình để PO theo dõi toàn cảnh dự án qua GitHub: scope, backlog, issues, milestones, wiki, risks, và tiến độ thực thi.

---

## 3. Business Objectives
1. Cung cấp một sản phẩm quản lý task đơn giản hơn Jira nhưng mạnh hơn các app todo cơ bản.
2. Cho phép self-host trên VPS với chi phí thấp.
3. Xây một kiến trúc đủ sạch để mở rộng lên:
   - Web production MVP
   - Mobile app (Flutter)
   - Desktop app (Electron)
   - SaaS multi-tenant roadmap
4. Thiết lập quy trình SDLC + Agile có khả năng audit và theo dõi đầy đủ trên GitHub.

---

## 4. Product Objectives
### MVP Objectives
- User có thể đăng ký, đăng nhập, đăng xuất.
- User có thể tạo, xem, sửa, xoá mềm task.
- Task có title, description, due date, priority, tags, assignee, status.
- Web UI responsive, có dark/light mode.
- Hệ thống có Docker, CI/CD, logging, health checks.

### Near-term Roadmap Objectives
- Monitoring & observability.
- Automated tests.
- Swagger / API documentation.
- Kanban / Gantt / reminders.
- Mobile + desktop planning-ready architecture.

---

## 5. Success Criteria
### Delivery Success
- Planning pack hoàn chỉnh và được PO chốt trước khi sang Phase 3 full execution.
- Tất cả scope MVP có issue, milestone, board status trên GitHub.
- Documentation structure đủ để dev mới onboard không cần hỏi lại context cơ bản.

### Product Success
- 5–20 users có thể dùng được trong nội bộ.
- API p95 < 200ms cho các endpoint chính ở tải MVP.
- Frontend load < 2 giây trong điều kiện bình thường.
- Uptime mục tiêu MVP: 99.5%.

---

## 6. Delivery Approach
- Methodology: Hybrid SDLC + Agile
- Planning Style: Engineering-ready, artifact-driven
- Tracking Tool: GitHub Issues, Milestones, Project Board, Wiki
- Source of Truth:
  1. GitHub Wiki
  2. docs/ in repo
  3. GitHub Project / Issues for execution state

---

## 7. SDLC Phases
### Phase 1 — Requirements & Planning
Deliverables:
- PRD
- Project Plan
- Scope
- WBS
- SRS
- Schedule / Milestones
- Budget / Resource Plan
- Risk Management Plan
- Use Cases
- Test Plan

### Phase 2 — System Design
Deliverables:
- Architecture docs
- SDD
- ERD / DB diagrams
- Sequence diagrams
- API contracts
- UX/UI system design

### Phase 3 — Implementation
Deliverables:
- Backend services
- Frontend app
- DB migrations
- CI/CD and deployment assets
- Feature-by-feature completion logs

### Phase 4 — Testing
Deliverables:
- Unit / integration / E2E suites
- Test report
- Defect log
- Release readiness checklist

### Phase 5 — Deployment
Deliverables:
- Production deploy guide
- Release notes
- Smoke test report
- Rollback procedure

### Phase 6 — Maintenance
Deliverables:
- Monitoring baseline
- Incident handling
- Backlog grooming
- Roadmap revision

---

## 8. Major Deliverables
| Deliverable | Owner | Status |
|---|---|---|
| Planning Pack | Tada + DangDDT | In Progress |
| PRD | Tada + PO | Drafted |
| SDD | Tada | Drafted |
| Backend MVP | Engineering | In Progress |
| Frontend MVP | Engineering | In Progress |
| Mobile Plan | Planning | Planned |
| Desktop Plan | Planning | Planned |
| DevOps Baseline | Engineering | In Progress |

---

## 9. Dependencies
- GitHub repo, wiki, project board available.
- VPS / environment details available before production deployment.
- PostgreSQL available for dev/test/prod.
- CI secrets configured before automated deployment.

---

## 10. Governance
- PO approves planning baseline.
- Tada maintains execution traceability.
- Every major artifact change should be reflected in GitHub docs and linked execution items.
- Every feature/bug/chore should map to issue(s) and project board state.

---

## 11. Exit Criteria for Planning Phase
Planning phase được xem là complete khi:
- Scope, assumptions, constraints được chốt.
- WBS và milestones rõ ràng.
- Risks được log và có mitigation.
- Use cases và SRS đủ để code.
- Test strategy đủ rõ để QA / dev thực thi.
- PO xác nhận planning baseline.
