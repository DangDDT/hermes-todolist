# Requirements Traceability Matrix (RTM)

## 1. Purpose
Requirements Traceability Matrix dùng để nối business intent với execution artifacts, giúp Product Owner theo dõi:
- yêu cầu nào đã được đặc tả chưa
- yêu cầu nào map vào use case nào
- endpoint nào phục vụ yêu cầu nào
- issue nào đang implement nó
- test case nào sẽ xác nhận nó

---

## 2. Traceability Rules
Mỗi requirement nên trace qua ít nhất các lớp sau:
1. PRD / Scope
2. SRS Functional Requirement
3. Use Case
4. API / UI / Component / Data Artifact
5. GitHub Issue / Task
6. Test Case / Test Suite

---

## 3. MVP RTM
| RTM ID | Business Need | PRD Ref | SRS Ref | Use Case | Design/Tech Artifact | GitHub Issue | Test Ref | Status |
|---|---|---|---|---|---|---|---|---|
| RTM-001 | User can register account | A.1 | FR-AUTH-01..03 | UC-01 | API Spec auth register, Backend Spec auth, Web Spec register form | #4, #9 | TC-AUTH-001, TC-AUTH-002, TS-BE-02, TS-FE-01 | Implemented |
| RTM-002 | User can login securely | A.2, A.3 | FR-AUTH-04..06 | UC-02 | API Spec auth login, middleware, Web Spec login | #4, #9 | TC-AUTH-003, TS-BE-02, TS-FE-01 | Implemented |
| RTM-003 | User can create task | B.1 | FR-TASK-01,02 | UC-03 | Backend Spec tasks, DB diagrams, Web Spec task form | #5, #12 | TC-TASK-001, TC-TASK-002, TS-BE-04, TS-FE-03 | Implemented |
| RTM-004 | User can view task list | B.2 | FR-TASK-03 | UC-04 | Web Spec list page, Sequence Diagrams list flow | #5, #11 | TC-TASK-003, TS-FE-02 | Implemented |
| RTM-005 | User can filter/sort tasks | B.3, B.4 | FR-TASK-04,05 | UC-04 | API Spec query params, Web Spec filter bar | #5, #11 | TC-TASK-003, TS-FE-02 | Implemented |
| RTM-006 | User can view task detail | B.2 | FR-TASK-06 | UC-05 | Web Spec detail page, Sequence Diagrams | #5, #13 | TS-FE-03 | Implemented |
| RTM-007 | User can update task | B.5 | FR-TASK-07 | UC-06 | Backend Spec task update, Web Spec edit page | #5, #12, #13 | TC-TASK-004, TS-BE-04, TS-FE-03 | Implemented |
| RTM-008 | User can soft delete task | B.6 | FR-TASK-08 | UC-07 | DB soft delete design, delete dialog, repository layer | #5, #11, #13 | TC-TASK-005, TS-BE-04, TS-FE-03 | Implemented |
| RTM-009 | UI has loading/error/empty states | C.3, C.4 | FR-UI-02,04 | UC-04..07 | Web Spec UX states, UX/UI Spec | #9, #11, #12, #13 | TC-UI-001, TC-UI-002, TS-FE-02/03 | Implemented |
| RTM-010 | System supports responsive themeable UI | C.1, C.2 | FR-UI-01,03 | Cross-cutting | UX/UI Spec, Web Spec, SDD | #8, #10 | TS-FE-02 (partial), visual QA | Implemented baseline |
| RTM-011 | System supports operational health checks | NFR Monitoring | FR-OPS-01 | UC-08 | DevOps Plan, Backend Spec | #25 | TC-OPS-001, TS-OPS-01 | Pending |
| RTM-012 | System supports structured logging | NFR Monitoring | FR-OPS-02 | UC-08 | DevOps Plan, Backend Spec | #25 | TS-OPS-01 | Partial |
| RTM-013 | System supports containerized deployment | NFR Deployment | FR-OPS-03 | Cross-cutting | DevOps Plan, docker docs | #14 | TS-OPS-01 | Implemented baseline |
| RTM-014 | System supports CI validation | NFR CI/CD | FR-OPS-04 | Cross-cutting | DevOps Plan, workflows | #15 | TS-OPS-01 | Implemented baseline |
| RTM-015 | System has backend automated test baseline | Quality goal | N/A | Cross-cutting | Test Plan | #16 | TS-BE-01..04 | Pending |
| RTM-016 | System has frontend automated test baseline | Quality goal | N/A | Cross-cutting | Test Plan | #17 | TS-FE-01..03 | Pending |
| RTM-017 | System has API documentation | Ops/Dev Enablement | N/A | Cross-cutting | API Spec, Swagger | #7 | Smoke/manual doc verification | Pending |

---

## 4. Roadmap Traceability Seeds
| RTM ID | Future Need | Backlog Ref | Candidate Artifacts |
|---|---|---|---|
| RTM-R01 | Kanban board view | BL-001 | Web Spec extension, new use cases, FE issues |
| RTM-R02 | Password reset | BL-004 | Auth SRS extension, sequence diagram, BE/FE issues |
| RTM-R03 | Mobile app | BL-017 | Mobile Spec, API stability matrix |
| RTM-R04 | Desktop app | BL-018 | Desktop Spec, packaging plan |
| RTM-R05 | Gantt / timeline view | BL-020 | UX/UI Spec extension, schedule-aware UI use cases |

---

## 5. Status Legend
- Implemented = feature baseline exists
- Partial = some technical baseline exists but acceptance not fully closed
- Pending = planned but not yet fully delivered
- Blocked = delivery blocked by dependency
