-- ============================================
-- Hermes TodoList — Migration 000001: Rollback
-- ============================================

DROP TRIGGER IF EXISTS trg_tasks_prevent_double_delete ON tasks;
DROP FUNCTION IF EXISTS prevent_double_delete();

DROP TRIGGER IF EXISTS trg_tasks_updated_at ON tasks;
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();

DROP TABLE IF EXISTS refresh_tokens CASCADE;
DROP TABLE IF EXISTS task_tags CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS task_priority;
DROP TYPE IF EXISTS task_status;
