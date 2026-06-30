1|## 2. Database Design
2|
3|### 2.1 ERD — Full Schema
4|
5|```mermaid
6|erDiagram
7|    users {
8|        uuid id PK "DEFAULT uuid_generate_v4()"
9|        varchar username UK "NOT NULL, 3-50 chars, regex: ^[a-z][a-z0-9_]*$"
10|        varchar password_hash "NOT NULL, bcrypt $2a$ cost=12"
11|        varchar display_name "NOT NULL, 1-100 chars, DEFAULT ''"
12|        timestamptz created_at "NOT NULL, DEFAULT NOW()"
13|        timestamptz updated_at "NOT NULL, DEFAULT NOW()"
14|    }
15|
16|    tasks {
17|        uuid id PK "DEFAULT uuid_generate_v4()"
18|        varchar title "NOT NULL, 1-255 chars"
19|        text description "NOT NULL, DEFAULT ''"
20|        task_status status "NOT NULL, DEFAULT 'TODO'"
21|        task_priority priority "NOT NULL, DEFAULT 'MEDIUM'"
22|        timestamptz due_date "NULLABLE"
23|        uuid creator_id FK "NOT NULL, ON DELETE CASCADE"
24|        uuid assignee_id FK "NULLABLE, ON DELETE SET NULL"
25|        timestamptz created_at "NOT NULL, DEFAULT NOW()"
26|        timestamptz updated_at "NOT NULL, DEFAULT NOW()"
27|        timestamptz deleted_at "NULLABLE, soft delete marker"
28|    }
29|
30|    tags {
31|        serial id PK
32|        varchar name UK "NOT NULL, 1-50 chars, lowercase"
33|        varchar color "NOT NULL, DEFAULT '#6B7280', hex format"
34|        timestamptz created_at "NOT NULL, DEFAULT NOW()"
35|    }
36|
37|    task_tags {
38|        uuid task_id PK_FK "ON DELETE CASCADE"
39|        int tag_id PK_FK "ON DELETE CASCADE"
40|    }
41|
42|    refresh_tokens {
43|        uuid id PK
44|        uuid user_id FK "NOT NULL, ON DELETE CASCADE"
45|        varchar token_hash "NOT NULL, sha256 of refresh token"
46|        timestamptz expires_at "NOT NULL"
47|        timestamptz created_at "NOT NULL, DEFAULT NOW()"
48|    }
49|
50|    users ||--o{ tasks : "creator (creator_id)"
51|    users ||--o{ tasks : "assignee (assignee_id)"
52|    tasks ||--o{ task_tags : "has"
53|    tags ||--o{ task_tags : "belongs to"
54|    users ||--o{ refresh_tokens : "owns"
55|```
56|
57|### 2.2 ENUM Definitions
58|
59|```sql
60|-- Task status lifecycle
61|CREATE TYPE task_status AS ENUM (
62|    'TODO',         -- Mới tạo, chưa bắt đầu
63|    'IN_PROGRESS',  -- Đang thực hiện
64|    'DONE',         -- Hoàn thành
65|    'CANCELLED'     -- Đã hủy
66|);
67|
68|-- Task priority levels
69|CREATE TYPE task_priority AS ENUM (
70|    'LOW',      -- Ưu tiên thấp
71|    'MEDIUM',   -- Ưu tiên trung bình (default)
72|    'HIGH',     -- Ưu tiên cao
73|    'URGENT'    -- Khẩn cấp
74|);
75|```
76|
77|### 2.3 State Machine — Task Lifecycle
78|
79|```mermaid
80|stateDiagram-v2
81|    [*] --> TODO : POST /tasks (create)
82|    
83|    TODO --> IN_PROGRESS : PATCH status=IN_PROGRESS
84|    TODO --> CANCELLED : PATCH status=CANCELLED
85|    
86|    IN_PROGRESS --> DONE : PATCH status=DONE
87|    IN_PROGRESS --> TODO : PATCH status=TODO (move back)
88|    IN_PROGRESS --> CANCELLED : PATCH status=CANCELLED
89|    
90|    DONE --> IN_PROGRESS : PATCH status=IN_PROGRESS (reopen)
91|    
92|    CANCELLED --> TODO : PATCH status=TODO (reactivate)
93|    
94|    DONE --> [*] : Soft delete (deleted_at)
95|    CANCELLED --> [*] : Soft delete (deleted_at)
96|    
97|    note right of DONE : Allowed transitions are<br/>validated at domain layer
98|```
99|
100|### 2.4 Index Strategy
101|
102|```sql
103|-- ============================================
104|-- Indexes — rationale & query patterns
105|-- ============================================
106|
107|-- 1. Task listing by creator (most frequent query)
108|-- Pattern: SELECT * FROM tasks WHERE creator_id = ? AND deleted_at IS NULL
109|CREATE INDEX idx_tasks_creator_active 
110|    ON tasks(creator_id) 
111|    WHERE deleted_at IS NULL;
112|
113|-- 2. Task listing by assignee (my tasks view)
114|-- Pattern: SELECT * FROM tasks WHERE assignee_id = ? AND deleted_at IS NULL
115|CREATE INDEX idx_tasks_assignee_active 
116|    ON tasks(assignee_id) 
117|    WHERE deleted_at IS NULL;
118|
119|-- 3. Filter by status (kanban columns)
120|-- Pattern: SELECT * FROM tasks WHERE status = ? AND deleted_at IS NULL
121|CREATE INDEX idx_tasks_status_active 
122|    ON tasks(status) 
123|    WHERE deleted_at IS NULL;
124|
125|-- 4. Sort by due date (upcoming tasks)
126|-- Pattern: SELECT * FROM tasks WHERE deleted_at IS NULL ORDER BY due_date
127|CREATE INDEX idx_tasks_due_date_active 
128|    ON tasks(due_date) 
129|    WHERE deleted_at IS NULL;
130|
131|-- 5. Sort by priority (priority view)
132|CREATE INDEX idx_tasks_priority_active 
133|    ON tasks(priority) 
134|    WHERE deleted_at IS NULL;
135|
136|-- 6. Full-text search on title (future)
137|-- Pattern: SELECT * FROM tasks WHERE to_tsvector('english', title) @@ to_tsquery(?)
138|-- CREATE INDEX idx_tasks_title_fts ON tasks USING GIN (to_tsvector('english', title));
139|
140|-- 7. Soft delete filtering
141|-- Ensures queries that include deleted records are still fast
142|CREATE INDEX idx_tasks_deleted 
143|    ON tasks(deleted_at);
144|
145|-- 8. Task-tag lookup
146|CREATE INDEX idx_task_tags_task 
147|    ON task_tags(task_id);
148|
149|-- 9. User lookup by username (login)
150|CREATE UNIQUE INDEX idx_users_username 
151|    ON users(username);
152|
153|-- 10. Refresh token lookup
154|CREATE INDEX idx_refresh_tokens_user 
155|    ON refresh_tokens(user_id);
156|CREATE INDEX idx_refresh_tokens_expires 
157|    ON refresh_tokens(expires_at);
158|```
159|
160|### 2.5 Triggers
161|
162|```sql
163|-- Auto-update updated_at on every row modification
164|CREATE OR REPLACE FUNCTION update_updated_at_column()
165|RETURNS TRIGGER AS $$
166|BEGIN
167|    NEW.updated_at = NOW();
168|    RETURN NEW;
169|END;
170|$$ LANGUAGE plpgsql;
171|
172|CREATE TRIGGER trg_users_updated_at
173|    BEFORE UPDATE ON users
174|    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
175|
176|CREATE TRIGGER trg_tasks_updated_at
177|    BEFORE UPDATE ON tasks
178|    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
179|
180|-- Prevent setting deleted_at on already-deleted tasks
181|CREATE OR REPLACE FUNCTION prevent_double_delete()
182|RETURNS TRIGGER AS $$
183|BEGIN
184|    IF OLD.deleted_at IS NOT NULL AND NEW.deleted_at IS NOT NULL THEN
185|        RAISE EXCEPTION 'Task % is already deleted', OLD.id;
186|    END IF;
187|    RETURN NEW;
188|END;
189|$$ LANGUAGE plpgsql;
190|
191|CREATE TRIGGER trg_tasks_prevent_double_delete
192|    BEFORE UPDATE ON tasks
193|    FOR EACH ROW
194|    WHEN (OLD.deleted_at IS NOT NULL)
195|    EXECUTE FUNCTION prevent_double_delete();
196|```
197|
198|### 2.6 Migration Files
199|
200|**000001_init.up.sql:** Full schema creation (users, tasks, tags, task_tags, refresh_tokens, enums, indexes, triggers, seed data)
201|
202|**000001_init.down.sql:** DROP all tables, types, functions in reverse order
203|
204|### 2.7 Seed Data
205|
206|```sql
207|-- Default tags (can be extended by admin)
208|INSERT INTO tags (name, color) VALUES
209|    ('bug',          '#EF4444'),  -- Red
210|    ('feature',      '#3B82F6'),  -- Blue
211|    ('improvement',  '#10B981'),  -- Green
212|    ('documentation','#F59E0B'),  -- Amber
213|    ('urgent',       '#8B5CF6'),  -- Purple
214|    ('design',       '#EC4899'),  -- Pink
215|    ('testing',      '#06B6D4'),  -- Cyan
216|    ('refactor',     '#F97316')   -- Orange
217|ON CONFLICT (name) DO NOTHING;
218|```
219|
220|### 2.8 Query Patterns (for sqlc)
221|
222|```sql
223|-- name: CreateUser :one
224|INSERT INTO users (username, password_hash, display_name) 
225|VALUES ($1, $2, $3) 
226|RETURNING *;
227|
228|-- name: GetUserByUsername :one
229|SELECT * FROM users WHERE username = $1;
230|
231|-- name: GetUserByID :one
232|SELECT * FROM users WHERE id = $1;
233|
234|-- name: CreateTask :one
235|INSERT INTO tasks (title, description, priority, due_date, creator_id, assignee_id) 
236|VALUES ($1, $2, $3, $4, $5, $6) 
237|RETURNING *;
238|
239|-- name: ListTasks :many
240|SELECT 
241|    t.*,
242|    u_creator.display_name AS creator_name,
243|    u_assignee.display_name AS assignee_name,
244|    COALESCE(
245|        json_agg(json_build_object('id', tg.id, 'name', tg.name, 'color', tg.color)) 
246|        FILTER (WHERE tg.id IS NOT NULL), '[]'
247|    ) AS tags
248|FROM tasks t
249|LEFT JOIN users u_creator ON t.creator_id = u_creator.id
250|LEFT JOIN users u_assignee ON t.assignee_id = u_assignee.id
251|LEFT JOIN task_tags tt ON t.id = tt.task_id
252|LEFT JOIN tags tg ON tt.tag_id = tg.id
253|WHERE t.deleted_at IS NULL
254|    AND (sqlc.narg('status')::task_status IS NULL OR t.status = sqlc.narg('status'))
255|    AND (sqlc.narg('assignee_id')::uuid IS NULL OR t.assignee_id = sqlc.narg('assignee_id'))
256|    AND (sqlc.narg('priority')::task_priority IS NULL OR t.priority = sqlc.narg('priority'))
257|    AND (sqlc.narg('tag')::varchar IS NULL OR tg.name = sqlc.narg('tag'))
258|GROUP BY t.id, u_creator.display_name, u_assignee.display_name
259|ORDER BY
260|    CASE WHEN sqlc.arg('sort_by') = 'due_date' AND sqlc.arg('order') = 'asc' 
261|         THEN t.due_date END ASC,
262|    CASE WHEN sqlc.arg('sort_by') = 'due_date' AND sqlc.arg('order') = 'desc' 
263|         THEN t.due_date END DESC,
264|    CASE WHEN sqlc.arg('sort_by') = 'priority' AND sqlc.arg('order') = 'asc' 
265|         THEN t.priority END ASC,
266|    CASE WHEN sqlc.arg('sort_by') = 'priority' AND sqlc.arg('order') = 'desc' 
267|         THEN t.priority END DESC,
268|    t.created_at DESC
269|LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
270|
271|-- name: GetTask :one
272|SELECT 
273|    t.*,
274|    u_creator.display_name AS creator_name,
275|    u_assignee.display_name AS assignee_name,
276|    COALESCE(
277|        json_agg(json_build_object('id', tg.id, 'name', tg.name, 'color', tg.color)) 
278|        FILTER (WHERE tg.id IS NOT NULL), '[]'
279|    ) AS tags
280|FROM tasks t
281|LEFT JOIN users u_creator ON t.creator_id = u_creator.id
282|LEFT JOIN users u_assignee ON t.assignee_id = u_assignee.id
283|LEFT JOIN task_tags tt ON t.id = tt.task_id
284|LEFT JOIN tags tg ON tt.tag_id = tg.id
285|WHERE t.id = $1 AND t.deleted_at IS NULL
286|GROUP BY t.id, u_creator.display_name, u_assignee.display_name;
287|
288|-- name: UpdateTask :one
289|UPDATE tasks SET
290|    title = COALESCE(sqlc.narg('title'), title),
291|    description = COALESCE(sqlc.narg('description'), description),
292|    status = COALESCE(sqlc.narg('status')::task_status, status),
293|    priority = COALESCE(sqlc.narg('priority')::task_priority, priority),
294|    due_date = COALESCE(sqlc.narg('due_date')::timestamptz, due_date),
295|    assignee_id = COALESCE(sqlc.narg('assignee_id')::uuid, assignee_id)
296|WHERE id = $1 AND deleted_at IS NULL
297|RETURNING *;
298|
299|-- name: SoftDeleteTask :exec
300|UPDATE tasks SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;
301|
302|-- name: ListTags :many
303|SELECT * FROM tags ORDER BY name;
304|
305|-- name: CreateRefreshToken :exec
306|INSERT INTO refresh_tokens (user_id, token_hash, expires_at) VALUES ($1, $2, $3);
307|
308|-- name: GetRefreshToken :one
309|SELECT * FROM refresh_tokens WHERE token_hash = $1 AND expires_at > NOW();
310|
311|-- name: DeleteRefreshToken :exec
312|DELETE FROM refresh_tokens WHERE token_hash = $1;
313|
314|-- name: DeleteUserRefreshTokens :exec
315|DELETE FROM refresh_tokens WHERE user_id = $1;
316|```
317|
318|### 2.9 Connection Pool Configuration
319|
320|```go
321|// Recommended pgx pool settings
322|poolConfig, _ := pgxpool.ParseConfig(databaseURL)
323|poolConfig.MaxConns = 25           // Maximum connections
324|poolConfig.MinConns = 5            // Minimum idle connections
325|poolConfig.MaxConnLifetime = 1h    // Max connection lifetime
326|poolConfig.MaxConnIdleTime = 30m   // Max idle time before close
327|poolConfig.HealthCheckPeriod = 1m  // Health check interval
328|```
329|
330|---
331|
332|---

## 3. Backend Design

### 3.1 Complete Package Tree

```
backend/
├── cmd/
│   ├── server/
│   │   └── main.go                  # Entry point: wire deps, start server
│   └── migrate/
│       └── main.go                  # Migration runner CLI
│
├── internal/
│   ├── config/
│   │   └── config.go                # Env-based config struct (caarlos0/env)
│   │
│   ├── domain/                      # Domain Layer (zero dependencies)
│   │   ├── task/
│   │   │   ├── entity.go            # Task aggregate root
│   │   │   ├── value_objects.go     # Status, Priority types + validation
│   │   │   └── repository.go        # TaskRepository interface (port)
│   │   └── user/
│   │       ├── entity.go            # User entity
│   │       ├── password.go          # bcrypt hash/compare
│   │       └── repository.go        # UserRepository interface (port)
│   │
│   ├── feature/                     # Feature Layer (use cases)
│   │   ├── auth_register/
│   │   ├── auth_login/
│   │   ├── auth_refresh/
│   │   ├── auth_logout/
│   │   ├── task_create/
│   │   ├── task_list/
│   │   ├── task_get/
│   │   ├── task_update/
│   │   ├── task_delete/
│   │   └── tag_list/
│   │
│   ├── infra/                       # Infrastructure Layer
│   │   ├── postgres/
│   │   │   ├── connection.go
│   │   │   ├── queries/             # sqlc generated
│   │   │   ├── sql/                 # SQL source files
│   │   │   └── migrations/
│   │   ├── repository/
│   │   │   ├── task_repo.go
│   │   │   └── user_repo.go
│   │   └── middleware/
│   │       ├── auth.go
│   │       ├── logger.go
│   │       ├── ratelimit.go
│   │       └── cors.go
│   │
│   └── shared/
│       ├── apperrors/
│       ├── response/
│       └── pagination/
│
└── docs/swagger/
```

### 3.2 Middleware Chain (Order Matters)

```
Request → RequestID → Logger → CORS → JWT Auth(opt) → Rate Limit → Recoverer → Handler
```

### 3.3 Error Handling Strategy

```go
type AppError struct {
    Code       string // "TASK_NOT_FOUND"
    Message    string
    HTTPStatus int
    Err        error  // wrapped original
}

// Constructors
func NotFound(code, message string) *AppError
func ValidationError(message string) *AppError
func Unauthorized(message string) *AppError
func Conflict(code, message string) *AppError
func Internal(err error) *AppError
```

### 3.4 Data Flow — Create Task

```
HTTP Request → Chi Router → Middleware Chain → Handler.ParseRequest
→ go-playground/validator → UseCase.Create()
  → Domain: task.New(title, desc, priority, dueDate) [validate business rules]
  → Domain: Task entity returned
  → Repository: Create(ctx, task) [sqlc INSERT]
  → DB returns row with ID + timestamps
→ Handler maps to JSON response
→ 201 Created + JSON body
```

---

## 4. API Design

### 4.1 Endpoint Map

| Method | Path | Auth | Feature |
|--------|------|------|---------|
| GET | /api/v1/health | No | Health check |
| POST | /api/v1/auth/register | No | Register |
| POST | /api/v1/auth/login | No | Login |
| POST | /api/v1/auth/refresh | Refresh Token | Token refresh |
| POST | /api/v1/auth/logout | Yes | Logout |
| GET | /api/v1/tasks | Yes | List tasks |
| POST | /api/v1/tasks | Yes | Create task |
| GET | /api/v1/tasks/{id} | Yes | Get task |
| PATCH | /api/v1/tasks/{id} | Yes | Update task |
| DELETE | /api/v1/tasks/{id} | Yes | Soft delete |
| GET | /api/v1/tags | No | List tags |

### 4.2 Standard Response Envelope

```json
// Success
{ "data": {...}, "meta": {"page":1, "per_page":20, "total":150} }

// Error
{ "error": { "code": "TASK_NOT_FOUND", "message": "...", "request_id": "req_abc" } }
```

### 4.3 Pagination

```
Request: ?page=2&per_page=20
Response: meta.page=2, meta.per_page=20, meta.total=150
Implementation: offset = (page-1)*per_page, limit = per_page (max 100)
```

### 4.4 Auth Token Flow

```
Login → bcrypt.Compare → Generate JWT (access 15min) + refresh token (7d)
→ Set-Cookie: access_token (httpOnly, Secure, SameSite=Strict)
→ Set-Cookie: refresh_token (httpOnly, path=/auth/refresh)
→ All API requests auto-include cookies

Refresh → Validate refresh_token hash in DB → Issue new pair → Rotate
```

### 4.5 Rate Limiting

| Zone | Rate | Burst | Scope |
|------|------|-------|-------|
| general | 30 req/s | 20 | All /api/* |
| auth | 5 req/s | 3 | /api/v1/auth/* |
| health | Unlimited | — | /api/v1/health |

---

## 5. Frontend Design

### 5.1 Route Map

```
/           → Redirect to /tasks
/login      → LoginPage (public)
/register   → RegisterPage (public)
/tasks      → TaskListPage (protected)
/tasks/new  → TaskCreatePage
/tasks/[id] → TaskDetailPage
/tasks/[id]/edit → TaskEditPage
```

### 5.2 Component Tree

```
<RootLayout>
  <Providers>  // TanStack Query + Auth + Theme
    <AuthGuard>
      // Public
      <LoginPage> → <LoginForm> → <FormField>×2 + <Button>
      <RegisterPage> → <RegisterForm> → <FormField>×3 + <Button>

      // Protected
      <DashboardLayout>
        <Sidebar> → <SidebarNav> + <ThemeToggle>
        <Header> → <Breadcrumb> + <UserMenu>
        <Main>
          <TaskListPage>
            <FilterBar> → Status, Priority, Assignee, Tag, Sort
            <TaskList> → <TaskCard>[] 
              <PriorityBadge> <StatusBadge> <TaskTags> <TaskAssignee> <TaskDueDate>
            <Pagination>
            <EmptyState>
          <TaskCreatePage> → <TaskForm mode="create">
          <TaskDetailPage>
            <TaskHeader> → <StatusBadge> <PriorityBadge> <EditButton> <DeleteButton>
            <TaskInfo> → <TaskDescription> <TaskMeta> <TaskTags>
          <TaskEditPage> → <TaskForm mode="edit">
```

### 5.3 State Management

```
AuthProvider (React Context)
  ├── user, isLoading, isAuthenticated
  ├── login(), register(), logout()

TanStack Query v5
  ├── useTasks(filters) → GET /tasks
  ├── useTask(id) → GET /tasks/:id
  ├── useCreateTask() → POST /tasks (optimistic update)
  ├── useUpdateTask() → PATCH /tasks/:id
  ├── useDeleteTask() → DELETE /tasks/:id
  └── useTags() → GET /tags (staleTime: 5min)

API Client (lib/api-client.ts)
  ├── Base URL from env
  ├── credentials: 'include' (cookies)
  ├── Auto token refresh on 401
  └── Error normalization
```

### 5.4 Zod Schemas

```typescript
export const LoginSchema = z.object({
  username: z.string().min(3).max(50),
  password: z.string().min(6).max(100),
});

export const CreateTaskSchema = z.object({
  title: z.string().min(1).max(255),
  description: z.string().optional().default(''),
  priority: z.enum(['LOW','MEDIUM','HIGH','URGENT']).optional().default('MEDIUM'),
  due_date: z.string().datetime().nullable().optional(),
  assignee_id: z.string().uuid().nullable().optional(),
  tag_ids: z.array(z.number().int().positive()).optional().default([]),
});
```

---

## 6. Mobile Design — Flutter

### 6.1 Architecture

```
lib/
├── main.dart              # Entry, ProviderScope
├── core/                  # Config, theme, router, network, utils
├── domain/                # Entities (freezed), repository interfaces
│   ├── task/              # Task entity + TaskRepository (abstract)
│   └── auth/              # User entity + AuthRepository
├── data/                  # Repository implementations (Dio)
│   ├── task/              # TaskRepositoryImpl + DTOs
│   └── auth/              # AuthRepositoryImpl + DTOs
└── presentation/          # Screens + Riverpod notifiers
    ├── auth/login/
    ├── auth/register/
    ├── tasks/list/
    ├── tasks/detail/
    ├── tasks/form/
    └── widgets/           # TaskCard, PriorityBadge, StatusBadge
```

### 6.2 Screen Map

| Route | Screen | Auth |
|-------|--------|------|
| /login | LoginScreen | No |
| /register | RegisterScreen | No |
| /tasks | TaskListScreen | Yes |
| /tasks/new | TaskFormScreen | Yes |
| /tasks/:id | TaskDetailScreen | Yes |
| /tasks/:id/edit | TaskFormScreen | Yes |

### 6.3 Key Decisions

| Decision | Choice | Why |
|----------|--------|-----|
| State | Riverpod 2.x | Compile-safe, testable |
| HTTP | Dio | Interceptors, retry |
| Models | freezed + json_serializable | Immutable |
| Routing | go_router | Deep linking |
| Secure Storage | flutter_secure_storage | Token storage |

### 6.4 Riverpod Providers

```dart
final authProvider = StateNotifierProvider<AuthNotifier, AuthState>(...);
final taskListProvider = FutureProvider.family<List<Task>, TaskFilters>(...);
final taskProvider = FutureProvider.family<Task, String>(...);
final createTaskProvider = StateNotifierProvider<TaskFormNotifier, ...>(...);
```

---

## 7. Desktop Design — Electron

### 7.1 Architecture

```
desktop/
├── src/main/              # Main process
│   ├── main.ts            # BrowserWindow, lifecycle
│   ├── tray.ts            # System tray
│   ├── shortcuts.ts       # Global shortcuts
│   ├── notifications.ts   # Native notifications
│   └── auto-updater.ts    # electron-updater
├── src/preload/
│   └── preload.ts         # contextBridge
└── src/renderer/          # = Next.js frontend (100% reuse)
```

### 7.2 Key Decisions

| Decision | Choice | Why |
|----------|--------|-----|
| Renderer | Reuse Next.js | DRY — 100% code reuse |
| Build | electron-builder | Cross-platform |
| Auto Update | electron-updater | GitHub Releases |
| IPC | contextBridge | Secure isolation |

### 7.3 Native Features

- System Tray: Quick-add task, show/hide, quit
- Global Shortcuts: Cmd/Ctrl+Shift+T for quick add
- Native Notifications: Due date reminders
- Window State: Remember position/size

---

## 8. DevOps Design

### 8.1 Docker Compose Topology

```
┌──────────┐   ┌──────────┐   ┌───────────────┐
│  nginx   │──→│   api    │──→│   postgres    │
│  :80:443 │   │  :8080   │   │   :5432       │
│  alpine  │   │  Go bin  │   │   16-alpine    │
└──────────┘   └──────────┘   └───────────────┘
     │                                
     └──────────→┌──────────┐         
                 │ frontend │         
                 │  :3000   │         
                 │ Next.js  │         
                 └──────────┘         
Network: hermes-network (bridge)
Volume: postgres_data
```

### 8.2 CI/CD

```
ci.yml (push, PR):
  Backend: lint (golangci-lint) → test (go test -race)
  Frontend: lint (ESLint) → type-check (tsc) → test (vitest)

deploy.yml (workflow_dispatch, tags v*):
  Build Docker images → Push GHCR → SSH VPS → docker compose pull → up -d
```

### 8.3 Monitoring

| Component | Tool |
|-----------|------|
| Health Check | GET /health + Docker HEALTHCHECK |
| Error Tracking | Sentry (Go: sentry-go, JS: @sentry/nextjs) |
| Logging | slog (JSON in prod) + Docker json-file driver |
| Metrics | (Future) Prometheus /metrics |

---

## 9. UX/UI Design System

### 9.1 Design Tokens

```css
/* Colors */
--color-primary: #3B82F6;     /* Actions, links */
--color-success: #10B981;     /* DONE status, success */
--color-warning: #F59E0B;    /* MEDIUM priority */
--color-danger:  #EF4444;    /* URGENT, delete, errors */

/* Typography */
--font-sans: 'Inter', sans-serif;
--font-mono: 'JetBrains Mono', monospace;
--text-xs: 0.75rem; --text-sm: 0.875rem; --text-base: 1rem;
--text-lg: 1.125rem; --text-xl: 1.25rem; --text-2xl: 1.5rem;

/* Spacing (4px base) */
--space-1: 4px; --space-2: 8px; --space-3: 12px; --space-4: 16px;
--space-6: 24px; --space-8: 32px;

/* Radius */
--radius-sm: 0.25rem; --radius-md: 0.5rem; --radius-lg: 0.75rem;

/* Shadows */
--shadow-sm: 0 1px 2px rgba(0,0,0,0.05);
--shadow-md: 0 4px 6px -1px rgba(0,0,0,0.1);
```

### 9.2 Status & Priority Colors

| Priority | Color | Status | Color |
|----------|-------|--------|-------|
| URGENT | #EF4444 Red | TODO | #6B7280 Gray |
| HIGH | #F97316 Orange | IN_PROGRESS | #3B82F6 Blue |
| MEDIUM | #F59E0B Amber | DONE | #10B981 Green |
| LOW | #6B7280 Gray | CANCELLED | #EF4444 Red |

### 9.3 Responsive Breakpoints

| Breakpoint | Width | Layout |
|-----------|-------|--------|
| Mobile | < 768px | Single column, hamburger menu |
| Tablet | 768-1024px | Collapsed sidebar (icons) |
| Desktop | > 1024px | Full sidebar + content |

---

## 10. Security Design

### 10.1 Threat Model (STRIDE)

| Threat | Mitigation |
|--------|-----------|
| Spoofing | JWT validation, bcrypt |
| Tampering | HTTPS, httpOnly cookies, input validation |
| Repudiation | Request IDs, audit log (backlog) |
| Info Disclosure | CORS, no stack traces in prod |
| DoS | Rate limiting (30 req/s), connection pooling |
| Elevation | JWT expiry, middleware guard |

### 10.2 Security Headers (Nginx)

```
X-Frame-Options: DENY
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Referrer-Policy: strict-origin-when-cross-origin
```

### 10.3 Secret Management

| Env | Storage |
|-----|---------|
| Local Dev | .env (gitignored) |
| CI/CD | GitHub Secrets |
| Production | Docker .env (chmod 600) |

---

## 11. Cross-Layer Consistency Check

### 11.1 Naming Convention

| Layer | Convention | Example |
|-------|-----------|---------|
| DB | snake_case | created_at |
| Go exported | PascalCase | TaskRepository |
| Go JSON | snake_case | json:"due_date" |
| TypeScript | camelCase | dueDate |
| HTTP Query | snake_case | ?sort_by=due_date |
| URL Path | kebab-case | /task-list |

### 11.2 Type Mapping

| PostgreSQL | Go | TypeScript | JSON |
|-----------|-----|-----------|------|
| UUID | uuid.UUID | string | "550e..." |
| VARCHAR(255) | string | string | "text" |
| TIMESTAMPTZ | time.Time | string (ISO) | "2026-..." |
| ENUM | Custom type | Union type | "TODO" |

### 11.3 Validation Consistency

| Field | DB | Go | Zod |
|-------|-----|-----|-----|
| username | VARCHAR(50) UNIQUE | min=3,max=50 | .min(3).max(50) |
| password | Not stored raw | min=6,max=100 | .min(6).max(100) |
| title | VARCHAR(255) | min=1,max=255 | .min(1).max(255) |
| priority | ENUM | oneof=LOW,MEDIUM,HIGH,URGENT | z.enum([...]) |

### 11.4 Layer Dependencies

```
Feature → depends on → Domain, Shared (NOT Infra)
Domain → NO deps
Infra → depends on → Domain (implements ports)
Shared → NO deps
```

---

## 12. Sign-off Checklist

- [ ] All C4 diagrams consistent
- [ ] DB schema matches domain entities
- [ ] API endpoints cover all features
- [ ] Frontend component tree complete
- [ ] Mobile mirrors web (backlog)
- [ ] Desktop reuses frontend 100%
- [ ] Docker topology matches C4
- [ ] CI/CD covers all quality gates
- [ ] Design tokens defined
- [ ] Security headers configured
- [ ] Naming consistent across layers
- [ ] Validation rules synchronized
- [ ] No circular dependencies

---

*System Design Document — v1.0 — Phase 2 Complete*  
*Next: Phase 3 — Implementation*

333|