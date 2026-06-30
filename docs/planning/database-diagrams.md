# Database Diagrams

## 1. MVP + Future-Ready ERD
```mermaid
erDiagram
    users ||--o{ tasks : creates
    users ||--o{ tasks : assigned_to
    tasks ||--o{ task_tags : has
    tags ||--o{ task_tags : maps
    users ||--o{ refresh_tokens : owns
    tasks ||--o{ comments : future
    tasks ||--o{ attachments : future
    tasks ||--o{ activity_logs : future

    users {
        uuid id
        varchar username
        varchar password_hash
        varchar display_name
        timestamptz created_at
        timestamptz updated_at
    }
    tasks {
        uuid id
        varchar title
        text description
        task_status status
        task_priority priority
        timestamptz due_date
        uuid creator_id
        uuid assignee_id
        timestamptz created_at
        timestamptz updated_at
        timestamptz deleted_at
    }
    tags {
        serial id
        varchar name
        varchar color
    }
    task_tags {
        uuid task_id
        int tag_id
    }
    refresh_tokens {
        uuid id
        uuid user_id
        varchar token_hash
        timestamptz expires_at
    }
    comments {
        uuid id
        uuid task_id
        uuid author_id
        text body
    }
    attachments {
        uuid id
        uuid task_id
        varchar storage_key
        varchar file_name
    }
    activity_logs {
        uuid id
        uuid task_id
        uuid actor_id
        varchar action
        jsonb metadata
    }
```

---

## 2. Data Ownership Notes
- `creator_id`: ai tạo task
- `assignee_id`: ai chịu trách nhiệm làm task
- `deleted_at`: soft delete boundary
- `refresh_tokens`: hỗ trợ auth hardening roadmap
- `comments`, `attachments`, `activity_logs`: future-ready entities, chưa thuộc MVP build scope
