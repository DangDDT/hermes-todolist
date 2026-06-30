-- ============================================
-- User Queries (sqlc)
-- ============================================

-- name: CreateUser :one
INSERT INTO users (username, password_hash, display_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, username, display_name, created_at
FROM users
ORDER BY display_name;
