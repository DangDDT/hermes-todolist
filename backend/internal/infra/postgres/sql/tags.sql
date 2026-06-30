-- ============================================
-- Tag Queries (sqlc)
-- ============================================

-- name: ListTags :many
SELECT * FROM tags ORDER BY name;

-- name: GetTagByName :one
SELECT * FROM tags WHERE name = $1;

-- name: CreateTag :one
INSERT INTO tags (name, color) VALUES ($1, $2) RETURNING *;
