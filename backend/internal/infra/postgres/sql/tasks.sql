-- ============================================
-- Task Queries (sqlc)
-- ============================================

-- name: CreateTask :one
INSERT INTO tasks (title, description, priority, due_date, creator_id, assignee_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListTasks :many
SELECT
    t.*,
    u_creator.display_name AS creator_name,
    u_assignee.display_name AS assignee_name
FROM tasks t
LEFT JOIN users u_creator ON t.creator_id = u_creator.id
LEFT JOIN users u_assignee ON t.assignee_id = u_assignee.id
WHERE t.deleted_at IS NULL
    AND (sqlc.narg('status')::task_status IS NULL OR t.status = sqlc.narg('status'))
    AND (sqlc.narg('assignee_id')::uuid IS NULL OR t.assignee_id = sqlc.narg('assignee_id'))
    AND (sqlc.narg('priority')::task_priority IS NULL OR t.priority = sqlc.narg('priority'))
    AND (sqlc.narg('creator_id')::uuid IS NULL OR t.creator_id = sqlc.narg('creator_id'))
ORDER BY
    CASE WHEN sqlc.arg('sort_by') = 'due_date' AND sqlc.arg('order') = 'asc'
         THEN t.due_date END ASC NULLS LAST,
    CASE WHEN sqlc.arg('sort_by') = 'due_date' AND sqlc.arg('order') = 'desc'
         THEN t.due_date END DESC NULLS LAST,
    CASE WHEN sqlc.arg('sort_by') = 'priority'
         THEN
            CASE t.priority
                WHEN 'URGENT' THEN 0
                WHEN 'HIGH' THEN 1
                WHEN 'MEDIUM' THEN 2
                WHEN 'LOW' THEN 3
            END
         END ASC,
    t.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountTasks :one
SELECT COUNT(*) FROM tasks t
WHERE t.deleted_at IS NULL
    AND (sqlc.narg('status')::task_status IS NULL OR t.status = sqlc.narg('status'))
    AND (sqlc.narg('assignee_id')::uuid IS NULL OR t.assignee_id = sqlc.narg('assignee_id'))
    AND (sqlc.narg('priority')::task_priority IS NULL OR t.priority = sqlc.narg('priority'))
    AND (sqlc.narg('creator_id')::uuid IS NULL OR t.creator_id = sqlc.narg('creator_id'));

-- name: GetTask :one
SELECT
    t.*,
    u_creator.display_name AS creator_name,
    u_assignee.display_name AS assignee_name
FROM tasks t
LEFT JOIN users u_creator ON t.creator_id = u_creator.id
LEFT JOIN users u_assignee ON t.assignee_id = u_assignee.id
WHERE t.id = $1 AND t.deleted_at IS NULL;

-- name: UpdateTask :one
UPDATE tasks SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    status = COALESCE(sqlc.narg('status')::task_status, status),
    priority = COALESCE(sqlc.narg('priority')::task_priority, priority),
    due_date = COALESCE(sqlc.narg('due_date')::timestamptz, due_date),
    assignee_id = COALESCE(sqlc.narg('assignee_id')::uuid, assignee_id)
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteTask :exec
UPDATE tasks SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: GetTaskTags :many
SELECT tg.* FROM tags tg
INNER JOIN task_tags tt ON tg.id = tt.tag_id
WHERE tt.task_id = $1
ORDER BY tg.name;

-- name: AddTaskTag :exec
INSERT INTO task_tags (task_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: RemoveTaskTag :exec
DELETE FROM task_tags WHERE task_id = $1 AND tag_id = $2;

-- name: SetTaskTags :exec
DELETE FROM task_tags WHERE task_id = $1;
-- Then call AddTaskTag for each tag in a loop
