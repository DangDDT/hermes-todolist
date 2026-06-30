# Test Plan / Test Cases / Test Suites

## 1. Test Strategy
Test approach: detailed, risk-based, layered.

### Test Levels
- Unit tests
- Integration tests
- API tests
- Component tests
- End-to-end tests
- Smoke tests

---

## 2. Scope
### In Scope
- Auth flows
- Task CRUD flows
- Validation and error handling
- Core UI states
- CI build integrity

### Out of Scope (initial baseline)
- Performance benchmark automation
- Full chaos/failure testing
- Multi-tenant security isolation

---

## 3. Test Suites
| Suite ID | Suite Name | Level |
|---|---|---|
| TS-BE-01 | Domain validation suite | Unit |
| TS-BE-02 | Auth usecase suite | Unit/Integration |
| TS-BE-03 | Task repository suite | Integration |
| TS-BE-04 | API endpoint suite | Integration |
| TS-FE-01 | Auth form validation suite | Component |
| TS-FE-02 | Task list rendering suite | Component |
| TS-FE-03 | Task detail/edit suite | Component |
| TS-E2E-01 | Critical user journey suite | E2E |
| TS-OPS-01 | Build & smoke suite | Smoke |

---

## 4. Sample Detailed Test Cases
### TC-AUTH-001 Register Success
- Preconditions: username not used
- Steps:
  1. Open register page
  2. Enter valid username, display name, password
  3. Submit form
- Expected:
  - API returns success
  - success feedback shown
  - user redirected to login or tasks depending on flow

### TC-AUTH-002 Register Duplicate Username
- Preconditions: username already exists
- Expected: conflict error shown, no account duplication

### TC-AUTH-003 Login Success
- Expected: authenticated session created, redirect to task list

### TC-TASK-001 Create Task Success
- Preconditions: authenticated user
- Expected: task persisted and visible in list

### TC-TASK-002 Create Task Missing Title
- Expected: validation error, no persistence

### TC-TASK-003 Filter by Status
- Expected: only matching tasks displayed

### TC-TASK-004 Update Task Priority
- Expected: saved value appears in detail/list

### TC-TASK-005 Soft Delete Task
- Expected: task removed from active list, record not hard deleted

### TC-UI-001 Loading State
- Expected: skeletons visible during fetch

### TC-UI-002 Empty State
- Expected: clear empty CTA shown when no tasks

### TC-OPS-001 Health Check
- Expected: `/api/v1/health` returns 200 and status payload

---

## 5. Release Gate
Không được release MVP nếu chưa có tối thiểu:
- Build pass backend/frontend
- Smoke tests pass
- Auth flow manual verification
- Task CRUD manual verification
- No open critical bugs
