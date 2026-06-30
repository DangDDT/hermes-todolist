# Backend Specification

## 1. Purpose
Define backend implementation contract for Hermes TodoList MVP and near-term expansion.

## 2. Architectural Style
- Language: Go
- Router: Chi
- Pattern: Clean Architecture + DDD-lite + feature-first delivery
- Persistence: PostgreSQL via pgx/sqlc-ready approach
- Interfaces:
  - HTTP handlers
  - Usecases
  - Domain entities/value objects
  - Repository layer

## 3. Module Boundaries
### Core Modules
- auth_register
- auth_login
- task_create
- task_list
- task_get
- task_update
- task_delete
- tag_list

### Shared Modules
- response envelope
- app errors
- pagination
- config
- logging
- auth middleware
- rate limit middleware

## 4. Domain Rules
### User
- username unique
- password hashed
- display_name defaults to username when omitted

### Task
- title required
- status lifecycle governed at domain layer
- priority constrained to enum values
- delete = soft delete only in MVP

## 5. Endpoint Matrix
| Endpoint | Purpose | Auth | Notes |
|---|---|---|---|
| POST /auth/register | register account | no | validate uniqueness |
| POST /auth/login | login | no | sets auth cookie |
| GET /tasks | list tasks | yes | pagination/filter/sort |
| POST /tasks | create task | yes | creator = current user |
| GET /tasks/{id} | task detail | yes | 404 if not found |
| PUT /tasks/{id} | full update | yes | validate transitions |
| DELETE /tasks/{id} | soft delete | yes | hide from active views |
| GET /tags | tag reference list | yes | may be cached later |

## 6. Validation Rules
- username: 3–50 chars, lowercase/alnum underscore
- password: min policy per auth spec
- title: 1–255 chars
- description: optional text
- due_date: ISO datetime or null
- priority: LOW/MEDIUM/HIGH/URGENT
- status: TODO/IN_PROGRESS/DONE/CANCELLED

## 7. Error Contract
Recommended JSON envelope:
- success response: `data`, optional `meta`
- error response: `error.code`, `error.message`, `error.details`

## 8. Security Controls
- bcrypt hashing
- JWT secret from config
- httpOnly cookies
- route protection middleware
- rate limiting
- input validation
- parameterized queries

## 9. Future Hooks
- refresh token rotation
- role-based access control
- audit logs
- comments/attachments/services
- background reminder jobs
