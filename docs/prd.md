# Product Requirements Document (PRD)

## v1.0 — MVP

**Product:** Hermes TodoList  
**Product Owner:** DangDDT  
**Date:** 2026-06-30  
**Status:** Draft  

---

## 1. Executive Summary

Hermes TodoList là ứng dụng quản lý công việc (todo list) hướng đến team nhỏ 5–20 người. MVP tập trung vào trải nghiệm đơn giản, nhanh, đáng tin cậy. Kiến trúc được thiết kế sẵn sàng mở rộng lên SaaS trong tương lai.

## 2. Problem Statement

Team nhỏ cần một công cụ quản lý task đơn giản, không cồng kềnh như Jira, không giới hạn như Todoist free tier, có thể self-host trên VPS riêng.

## 3. Target Users

| Persona | Mô tả | Nhu cầu chính |
|---------|-------|--------------|
| **Team Member** | Thành viên trong team 5-20 người | Xem task được giao, cập nhật trạng thái |
| **Team Lead** | Quản lý nhóm | Giao task, theo dõi tiến độ, set priority |
| **Admin** | Quản trị hệ thống | Quản lý users, cấu hình hệ thống |

## 4. MVP Scope

### 4.1 MUST HAVE (MVP v1.0)

#### A. Authentication
- **A.1** Đăng ký tài khoản với username + password
- **A.2** Đăng nhập / đăng xuất
- **A.3** JWT-based session management
- **A.4** Password hashing (bcrypt/argon2)

#### B. Task Management
- **B.1** Tạo task mới với các trường:
  - Title (bắt buộc)
  - Description (tùy chọn)
  - Due date (tùy chọn)
  - Priority: LOW, MEDIUM, HIGH, URGENT
  - Tags (nhiều tags/task)
  - Assignee (1 người)
  - Status: TODO, IN_PROGRESS, DONE, CANCELLED
- **B.2** Xem danh sách task (list view)
- **B.3** Lọc task theo: status, priority, assignee, tags
- **B.4** Sắp xếp theo: due date, priority, created date
- **B.5** Cập nhật task (edit từng trường)
- **B.6** Xoá task (soft delete)

#### C. UI/UX
- **C.1** Responsive web design
- **C.2** Dark mode / Light mode
- **C.3** Loading states, empty states, error states
- **C.4** Toast notifications cho hành động thành công/thất bại

### 4.2 NICE TO HAVE (Backlog)

- **D.1** Mobile app (Flutter)
- **D.2** Desktop app (Electron)
- **D.3** Real-time updates (WebSocket)
- **D.4** Kanban board view
- **D.5** Subtasks / task nesting
- **D.6** Comments trên task
- **D.7** File attachments
- **D.8** Activity log / audit trail
- **D.9** AI integration (Hermes Agent)
- **D.10** Multi-tenancy cho SaaS
- **D.11** OAuth login (Google, GitHub)
- **D.12** Email notifications

## 5. Non-Functional Requirements

| Category | Requirement |
|----------|------------|
| **Performance** | API response < 200ms p95, page load < 2s |
| **Scalability** | Hỗ trợ 5-20 concurrent users (MVP), thiết kế scale được lên 1000+ |
| **Security** | JWT auth, password hashing, input validation, SQL injection prevention, CORS, rate limiting |
| **Reliability** | Uptime 99.5% (MVP), graceful error handling |
| **Deployment** | Docker Compose, single VPS (MVP) |
| **CI/CD** | GitHub Actions: lint → test → build → deploy |
| **Monitoring** | Health check endpoint, structured logging, error tracking (Sentry), basic metrics |

## 6. Constraints & Assumptions

### Constraints
- MVP chỉ 1 developer (DangDDT)
- VPS single node (không Kubernetes)
- Ngân sách tối thiểu (~$10-20/tháng VPS)

### Assumptions
- Users có trình duyệt hiện đại (Chrome, Firefox, Safari, Edge)
- Team nội bộ, không cần multi-tenancy ngay
- Không cần offline mode cho MVP

## 7. Success Metrics (MVP)

| Metric | Target |
|--------|--------|
| Users đăng ký | ≥ 5 (internal team) |
| Tasks created | ≥ 50 trong tháng đầu |
| Crash-free rate | ≥ 99% |
| Page load time | < 2 giây |
| User retention | ≥ 80% sau 1 tháng |

## 8. Glossary

| Term | Definition |
|------|-----------|
| MVP | Minimum Viable Product |
| PRD | Product Requirements Document |
| SDLC | Software Development Life Cycle |
| VPS | Virtual Private Server |
| JWT | JSON Web Token |

---

*Document maintained by Tada (Hermes Agent) as part of SDLC Phase 1.*
