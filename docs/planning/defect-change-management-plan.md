# Defect & Change Management Plan

## 1. Purpose
Tài liệu này chuẩn hóa cách quản lý bug/defect và cách xử lý yêu cầu thay đổi scope trong quá trình delivery.

---

## 2. Defect Management
### 2.1 Defect Lifecycle
1. Defect discovered
2. Defect logged on GitHub
3. Triage severity + priority
4. Assigned to backlog / current sprint / hotfix lane
5. Fixed
6. Verified
7. Closed
8. Reopened if verification fails

### 2.2 Defect Issue Format
Recommended title:
- `[BUG] <component> — <short defect description>`

Required fields:
- Summary
- Environment / context
- Steps to reproduce
- Expected behavior
- Actual behavior
- Severity
- Priority
- Root cause (when known)
- Fix summary
- Verification notes

### 2.3 Severity Levels
| Severity | Meaning |
|---|---|
| Critical | system down, security breach, data loss risk |
| High | core feature blocked or broken |
| Medium | feature degraded but workaround exists |
| Low | cosmetic/minor usability issue |

### 2.4 Priority Levels
| Priority | Meaning |
|---|---|
| P0 | immediate hotfix |
| P1 | current sprint / release blocking |
| P2 | next sprint |
| P3 | backlog |

### 2.5 Reopen Rules
Issue should be reopened if:
- expected fix not fully works
- regression found in same acceptance boundary
- root cause not actually resolved

---

## 3. Change Management
### 3.1 What Counts as Change Request
A change request is any modification that affects at least one of:
- scope
- acceptance criteria
- timeline/milestone
- architecture baseline
- budget/resource assumptions
- risk profile

### 3.2 Change Request Lifecycle
1. Change proposed
2. Impact assessed
3. PO decision
4. Baseline docs updated
5. Backlog/issues updated
6. Schedule/risk/test implications logged

### 3.3 Change Categories
| Category | Example |
|---|---|
| Scope Add | add comments feature into MVP |
| Scope Reduce | postpone reminders to backlog |
| Technical Change | switch auth/session strategy |
| Schedule Change | move milestone due date |
| Quality Change | raise release gate coverage threshold |

### 3.4 Impact Assessment Checklist
For each change ask:
- Which docs need update?
- Which issues are affected?
- Does API contract change?
- Does database schema change?
- Does UI flow change?
- Does test scope change?
- Does milestone shift?
- Does risk increase?

### 3.5 Approval Rules
| Change Type | Approver |
|---|---|
| business scope change | PO |
| architecture-impacting change | PO + technical consultation |
| release timeline change | PO |
| minor internal implementation adjustment | execution owner, PO informed if user-visible impact exists |

---

## 4. GitHub Operational Rules
### For Defects
- create `type: bug`
- assign severity/priority labels
- link affected feature issue if applicable
- move project board state accordingly

### For Changes
- create `type: chore` or `type: docs` change record if documentation baseline affected
- update milestone if schedule changes
- update wiki/repo docs before or with implementation

---

## 5. Definition of Done for Bug Fix
A bug fix is done only when:
- root cause identified or at least plausible cause documented
- fix implemented
- build/tests pass as applicable
- verification performed
- issue updated with fix summary
- related docs updated if behavior changed

---

## 6. Definition of Done for Approved Change
A change is done only when:
- PO decision captured
- impacted artifacts updated
- impacted issues/tasks updated
- project board state consistent
- traceability maintained
