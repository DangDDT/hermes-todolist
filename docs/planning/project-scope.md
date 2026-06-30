# Project Scope — Hermes TodoList

## 1. Scope Statement
Hermes TodoList MVP bao gồm web application cho quản lý task nội bộ team nhỏ, cùng backend API, database, deployment baseline, và quy trình vận hành dự án trên GitHub.

---

## 2. In-Scope — MVP v1.0
### Product Scope
#### Authentication
- Username/password registration
- Login/logout
- JWT cookie-based session
- Password hashing

#### Task Management
- Create task
- View list
- View detail
- Edit task
- Soft delete task
- Filter by status, priority, assignee, tags
- Sort by due date, priority, created date

#### UX/UI
- Responsive web app
- Dashboard shell
- Toast / loading / empty / error states
- Light/dark mode

#### DevOps / Operational Baseline
- Docker Compose
- Reverse proxy baseline
- CI pipeline
- Build verification
- Health check endpoint
- Structured logging

#### Documentation / Governance
- Wiki + docs structure
- GitHub issues / project / milestones
- Planning artifacts
- Backlog and roadmap tracking

---

## 3. Near-Term Scope (Roadmap gần)
- Swagger docs
- Monitoring & observability
- Backend/Frontend test suites
- Password reset
- Kanban view
- Gantt / timeline view
- Reminder jobs
- Comments / activity logs
- Mobile app design baseline
- Desktop app design baseline

---

## 4. Out of Scope — MVP
- Multi-tenancy
- Billing / subscriptions
- Marketplace integrations
- SSO / SAML
- Full RBAC matrix beyond MVP-level roles
- Offline-first sync
- Real-time collaboration
- Enterprise audit/compliance module

---

## 5. Scope by Platform
| Platform | Scope Level |
|---|---|
| Web | Full MVP implementation |
| Backend API | Full MVP implementation |
| Database | Full MVP schema with future-ready hooks |
| Mobile | Planning + roadmap, not full build in current planning baseline |
| Desktop | Planning + roadmap, not full build in current planning baseline |
| DevOps | MVP production baseline |

---

## 6. Scope Constraints
- Team assumption: 3–5 contributors equivalent.
- MVP infra: single VPS.
- Low-cost hosting target.
- No Kubernetes in MVP.
- No multi-region or HA architecture in MVP.

---

## 7. Acceptance Boundary
Nếu một feature không trực tiếp phục vụ khả năng:
1. auth,
2. task CRUD,
3. usable responsive UI,
4. deployable MVP,
5. observable baseline,

thì mặc định nó thuộc roadmap/backlog, không tự động nằm trong MVP scope.
