-- ============================================
-- Hermes TodoList — Migration 000002: Task comments
-- ============================================

CREATE TABLE task_comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_task_comments_task_created_at ON task_comments(task_id, created_at DESC);
CREATE INDEX idx_task_comments_author ON task_comments(author_id);
