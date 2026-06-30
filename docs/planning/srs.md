# Software Requirements Specification (SRS)

## 1. Purpose
Tài liệu này đặc tả yêu cầu hệ thống cho Hermes TodoList ở mức đủ chi tiết để design, implement, test và trace ngược về business objective.

## 2. Product Scope
Hermes TodoList là hệ thống quản lý task cho team nhỏ, hỗ trợ xác thực người dùng, CRUD task, tracking trạng thái, và một web UI responsive có khả năng deploy self-host.

## 3. Stakeholders
- Product Owner
- Team Member/User
- Admin
- Engineering
- Future Support/Operations

## 4. System Features
### 4.1 User Authentication
Functional requirements:
- FR-AUTH-01: System shall allow user registration with username + password.
- FR-AUTH-02: System shall validate unique username.
- FR-AUTH-03: System shall hash password before persistence.
- FR-AUTH-04: System shall authenticate valid login requests.
- FR-AUTH-05: System shall issue JWT-based authenticated session.
- FR-AUTH-06: System shall reject invalid credentials.

### 4.2 Task Lifecycle Management
- FR-TASK-01: System shall allow authenticated user to create task.
- FR-TASK-02: System shall persist title, description, due date, priority, assignee, status.
- FR-TASK-03: System shall allow list retrieval with pagination.
- FR-TASK-04: System shall support filtering by status, priority, assignee, tags.
- FR-TASK-05: System shall support sorting by due date, priority, created date.
- FR-TASK-06: System shall allow viewing task detail.
- FR-TASK-07: System shall allow updating task fields.
- FR-TASK-08: System shall support soft delete.

### 4.3 UI / Experience
- FR-UI-01: System shall provide responsive web layout.
- FR-UI-02: System shall provide loading/empty/error states.
- FR-UI-03: System shall provide dark/light theme toggle.
- FR-UI-04: System shall provide success/error feedback for user actions.

### 4.4 Operational Support
- FR-OPS-01: System shall expose health check endpoint.
- FR-OPS-02: System shall produce structured logs.
- FR-OPS-03: System shall support containerized deployment.
- FR-OPS-04: System shall provide CI validation on build pipeline.

## 5. External Interface Requirements
### 5.1 User Interface
- Browser-based UI
- Dashboard shell with navigation
- Forms for auth and task management

### 5.2 Software Interfaces
- PostgreSQL
- GitHub Actions
- Nginx reverse proxy
- Future: Sentry, Prometheus/Grafana, email provider

### 5.3 API Interfaces
- RESTful JSON API under `/api/v1`
- Cookie-based auth session

## 6. Non-Functional Requirements
### Performance
- NFR-PERF-01: p95 API response < 200ms for core endpoints under MVP load.
- NFR-PERF-02: Initial page render < 2s in target environment.

### Reliability
- NFR-REL-01: System shall gracefully handle recoverable errors.
- NFR-REL-02: Health endpoint shall indicate service health.

### Security
- NFR-SEC-01: Passwords shall never be stored in plaintext.
- NFR-SEC-02: Input shall be validated server-side.
- NFR-SEC-03: Sensitive endpoints shall require authenticated access.
- NFR-SEC-04: System shall mitigate SQL injection through parameterized queries.
- NFR-SEC-05: System shall support rate limiting.

### Maintainability
- NFR-MAINT-01: Backend shall use layered architecture.
- NFR-MAINT-02: Frontend shall use feature-oriented organization.
- NFR-MAINT-03: Source of truth artifacts shall remain version-controlled.

## 7. Assumptions
- Team size 3–5 equivalent contributors.
- Single VPS deployment for MVP.
- Users have modern browsers.

## 8. Constraints
- No Kubernetes in MVP.
- Budget-conscious infra.
- MVP timelines prioritize core functionality over enterprise features.

## 9. Traceability Summary
| Requirement Group | Primary Docs |
|---|---|
| Auth | PRD, API Spec, SDD, Use Cases |
| Task CRUD | PRD, API Spec, SDD, Test Plan |
| DevOps | Project Plan, Schedule, Risk Plan |
| UX/UI | SDD, Use Cases, Test Plan |
