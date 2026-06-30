# Backlog

> Danh sách tính năng & ý tưởng phát triển tương lai.  
> Prioritized by business value & technical feasibility. PO (DangDDT) quyết định priority cuối cùng.

---

## 🗂️ Feature Backlog

### P1 — Phase 2 (Sau MVP ổn định)

| ID | Feature | Effort | Notes |
|----|---------|--------|-------|
| BL-001 | Kanban Board View | M | Thay thế/toggle với List View |
| BL-002 | Subtasks / Task Nesting | M | Parent-child relationship |
| BL-003 | Due Date Reminders (Email) | S | Cron job check due dates |
| BL-004 | Password Reset Flow | S | Email-based reset token |

### P2 — Phase 3 (Mở rộng team)

| ID | Feature | Effort | Notes |
|----|---------|--------|-------|
| BL-005 | Real-time Updates (WebSocket) | L | Live sync giữa các user |
| BL-006 | Comments on Tasks | M | Thread-based discussion |
| BL-007 | Activity Log / Audit Trail | M | Ai làm gì, lúc nào |
| BL-008 | File Attachments | M | Upload + S3/MinIO storage |
| BL-009 | Advanced Search & Filter | M | Full-text search, saved filters |
| BL-010 | Task Templates | S | Reusable task presets |
| BL-011 | Team Roles (Admin/Member/Viewer) | M | RBAC authorization |

### P3 — Phase 4 (SaaS Launch)

| ID | Feature | Effort | Notes |
|----|---------|--------|-------|
| BL-012 | Multi-tenancy | L | Workspace/Organization isolation |
| BL-013 | OAuth Login (Google, GitHub) | M | Social authentication |
| BL-014 | Billing & Subscription | L | Stripe integration |
| BL-015 | API Rate Limiting & Quotas | M | Per-tenant limits |
| BL-016 | Public REST API | L | API keys, documentation portal |
| BL-017 | Mobile App (Flutter) | XL | iOS + Android |
| BL-018 | Desktop App (Electron) | L | macOS + Windows + Linux |

### P4 — Future / Moon Shot

| ID | Feature | Effort | Notes |
|----|---------|--------|-------|
| BL-019 | AI Integration (Hermes Agent) | L | Voice commands, auto-prioritize |
| BL-020 | Gantt Chart / Timeline View | L | Project planning view |
| BL-021 | Time Tracking | M | Pomodoro, time logs |
| BL-022 | Integration Marketplace | XL | Webhooks, Zapier, n8n |
| BL-023 | Offline-First Mode | L | Local-first with sync |
| BL-024 | SSO / SAML (Enterprise) | M | Okta, Azure AD |

---

## 🐛 Technical Debt & Improvements

| ID | Item | Effort | Notes |
|----|------|--------|-------|
| TD-001 | API Pagination Cursor-based | M | Better perf than offset |
| TD-002 | Redis Caching Layer | M | Reduce DB load |
| TD-003 | Database Read Replicas | L | Scale read operations |
| TD-004 | API Versioning Strategy | S | URL-based vs header-based |
| TD-005 | E2E Test Suite (Playwright) | L | Critical path coverage |
| TD-006 | Load Testing (k6) | M | Performance baseline |
| TD-007 | Grafana + Prometheus Dashboard | M | Advanced monitoring |
| TD-008 | Database Backup Strategy | M | Automated pg_dump + retention |
| TD-009 | Infrastructure as Code (Terraform) | L | Reproducible infra |
| TD-010 | Blue-Green Deployment | M | Zero-downtime deploys |

---

## 📊 Summary

| Phase | Features | Tech Debt | Total Items |
|-------|----------|-----------|-------------|
| P1 (Phase 2) | 4 | 0 | 4 |
| P2 (Phase 3) | 7 | 0 | 7 |
| P3 (Phase 4) | 7 | 0 | 7 |
| P4 (Future) | 6 | 0 | 6 |
| Tech Debt | 0 | 10 | 10 |
| **TOTAL** | **24** | **10** | **34** |

---

## 🔄 Backlog Process

1. PO (DangDDT) hoặc Tada đề xuất item mới
2. PO đánh giá priority (P1-P4)
3. Ghi vào đây kèm effort estimate
4. Khi bắt đầu phase mới → PO pick items → Tada plan + implement

---

*Last updated: 2026-06-30 by Tada*
