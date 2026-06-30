# Web Frontend Specification

## 1. Purpose
Define the MVP web experience and interaction contract.

## 2. Stack
- Next.js App Router
- TypeScript strict
- TanStack Query
- react-hook-form + zod
- shadcn/ui v4
- Tailwind CSS v4

## 3. Route Map
| Route | Purpose | Access |
|---|---|---|
| /login | login form | public |
| /register | register form | public |
| /tasks | task list | authenticated |
| /tasks/new | create task | authenticated |
| /tasks/{id} | detail | authenticated |
| /tasks/{id}/edit | edit | authenticated |
| /settings | settings placeholder | authenticated |

## 4. UX States Required
Every data-driven screen must have:
- loading state
- empty state
- error state
- success feedback
- navigation recovery path

## 5. Task List Requirements
- show title, status, priority, due date, assignee, tags
- filter by status, priority, search
- sort by due date, priority, created date
- paginate results
- CTA to create task

## 6. Forms
### Register
- username
- display name
- password
- confirm password

### Login
- username
- password

### Task Create/Edit
- title
- description
- priority
- status (edit only or controlled flow)
- due date
- assignee
- tags (future enhancement if API not fully ready)

## 7. Accessibility Baseline
- keyboard navigable forms
- visible focus states
- semantic labels
- sufficient contrast
- screen-reader friendly text on icon buttons

## 8. Future Expansion Hooks
- Kanban board
- timeline view
- comments panel
- activity drawer
- notifications center
