# 🚀 Hermes TodoList

> **Production-grade todo list app** — Full SDLC · Team 5-20 users · SaaS-ready architecture

<p align="center">

[![CI](https://github.com/DangDDT/hermes-todolist/actions/workflows/ci.yml/badge.svg)](https://github.com/DangDDT/hermes-todolist/actions/workflows/ci.yml)
[![Deploy](https://github.com/DangDDT/hermes-todolist/actions/workflows/deploy.yml/badge.svg)](https://github.com/DangDDT/hermes-todolist/actions/workflows/deploy.yml)
[![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8?logo=go)](https://go.dev)
[![Next.js](https://img.shields.io/badge/next.js-15+-black?logo=next.js)](https://nextjs.org)
[![PostgreSQL](https://img.shields.io/badge/postgres-16-316192?logo=postgresql)](https://www.postgresql.org)
[![Docker](https://img.shields.io/badge/docker-ready-2496ED?logo=docker)](https://docker.com)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

</p>

---

## 📊 Project Dashboard

| Dashboard | Link |
|-----------|------|
| 🗂️ **Issues** | [View All](https://github.com/DangDDT/hermes-todolist/issues) |
| 🏷️ **Labels** | [Filter by label](https://github.com/DangDDT/hermes-todolist/labels) |
| 🎯 **Milestones** | [MVP v1.0](https://github.com/DangDDT/hermes-todolist/milestone/1) · [P1](https://github.com/DangDDT/hermes-todolist/milestone/2) · [P2](https://github.com/DangDDT/hermes-todolist/milestone/3) · [P3](https://github.com/DangDDT/hermes-todolist/milestone/4) · [P4](https://github.com/DangDDT/hermes-todolist/milestone/5) |
| 🔄 **CI/CD** | [Actions](https://github.com/DangDDT/hermes-todolist/actions) |
| 📚 **Wiki** | [Documentation Hub](https://github.com/DangDDT/hermes-todolist/wiki) |
| 🛡️ **Security** | [Security Advisories](https://github.com/DangDDT/hermes-todolist/security) |
| 📈 **Insights** | [Pulse](https://github.com/DangDDT/hermes-todolist/pulse) · [Contributors](https://github.com/DangDDT/hermes-todolist/graphs/contributors) |

---

## 🎯 MVP v1.0 Progress

```mermaid
gantt
    title SDLC Progress
    dateFormat  YYYY-MM-DD
    axisFormat  %b %d

    section 📋 Phase 1: Planning
    PRD & Requirements     :done,    p1, 2026-06-30, 1d

    section 🏗️ Phase 2: Design
    Backend (Go + Chi)     :done,    p2a, 2026-06-30, 1d
    Frontend (Next.js)     :done,    p2b, 2026-06-30, 1d
    Database (PostgreSQL)  :done,    p2c, 2026-06-30, 1d
    DevOps (Docker + CI/CD):active,  p2d, 2026-07-01, 2d
    UX/UI (Design System)  :         p2e, 2026-07-02, 2d

    section 💻 Phase 3: Implementation
    Backend Scaffold       :         p3a, 2026-07-04, 2d
    Auth Module            :         p3b, 2026-07-06, 2d
    Task CRUD              :         p3c, 2026-07-08, 3d
    Frontend Scaffold      :         p3d, 2026-07-04, 2d
    Auth Pages             :         p3e, 2026-07-06, 2d
    Task Pages             :         p3f, 2026-07-08, 3d

    section 🧪 Phase 4: Testing
    Backend Tests          :         p4a, 2026-07-11, 2d
    Frontend Tests         :         p4b, 2026-07-11, 2d
    Integration Tests      :         p4c, 2026-07-13, 2d

    section 🚀 Phase 5: Deploy
    Docker Setup           :         p5a, 2026-07-15, 1d
    VPS Deploy             :         p5b, 2026-07-16, 1d

    section 🔧 Phase 6: Maintain
    Monitoring & Ops       :         p6, 2026-07-17, 14d
```

---

## 🏗️ Architecture

```mermaid
graph TB
    subgraph Client["🖥️ Client"]
        Web["Next.js Web App<br/>:3000"]
    end

    subgraph Gateway["🚪 Gateway"]
        Nginx["Nginx :80/:443"]
    end

    subgraph Backend["⚙️ API"]
        API["Go + Chi<br/>:8080"]
    end

    subgraph Data["💾 Data"]
        DB["PostgreSQL 16<br/>:5432"]
    end

    Web --> Nginx --> API --> DB

    style Client fill:#1a1a2e,stroke:#3b82f6,color:#fff
    style Gateway fill:#1a1a2e,stroke:#f59e0b,color:#fff
    style Backend fill:#1a1a2e,stroke:#10b981,color:#fff
    style Data fill:#1a1a2e,stroke:#ef4444,color:#fff
```

👉 [Full Architecture Docs](https://github.com/DangDDT/hermes-todolist/wiki/Architecture)

---

## 🛠️ Tech Stack

| Layer | Stack |
|-------|-------|
| **Frontend** | Next.js 15 (App Router) · shadcn/ui · TanStack Query v5 · Tailwind v4 · react-hook-form + zod |
| **Backend** | Go 1.23+ · Chi v5 · sqlc · golang-migrate · swaggo/swag · slog · caarlos0/env |
| **Database** | PostgreSQL 16 · UUID PK · Soft Delete · Partial Indexes |
| **DevOps** | Docker Compose · GitHub Actions · Nginx · GHCR |

👉 [Full Tech Stack Docs](https://github.com/DangDDT/hermes-todolist/wiki/Tech-Stack)

---

## 📋 MVP Features

- [x] Project documentation & planning
- [ ] User authentication (username + password)
- [ ] Task CRUD (title, description, due date, priority, tags, assignee, status)
- [ ] Task list with filtering, sorting, pagination
- [ ] Responsive web UI with dark mode
- [ ] REST API with Swagger docs
- [ ] Docker Compose deployment
- [ ] CI/CD pipeline
- [ ] Unit + integration tests

👉 [Full PRD](https://github.com/DangDDT/hermes-todolist/wiki/PRD) · [Backlog (34 items)](https://github.com/DangDDT/hermes-todolist/wiki/Backlog)

---

## 🏃 Quick Start

```bash
# Clone
git clone https://github.com/DangDDT/hermes-todolist.git
cd hermes-todolist

# Start all services
docker compose up -d

# API at http://localhost:8080
# Swagger at http://localhost:8080/swagger/
# Web at http://localhost:3000
```

---

## 📂 Project Structure

```
hermes-todolist/
├── .github/
│   ├── workflows/
│   │   ├── ci.yml              # Lint + Test on push/PR
│   │   └── deploy.yml          # Build Docker + Deploy to VPS
│   └── dependabot.yml          # Auto dependency updates
├── backend/                    # Go API
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── config/
│   │   ├── domain/             # DDD Entities
│   │   ├── feature/            # Feature-first handlers
│   │   ├── infra/              # PostgreSQL, middleware
│   │   └── shared/             # Errors, response, pagination
│   ├── Dockerfile
│   └── Makefile
├── frontend/                   # Next.js web app
│   ├── src/
│   │   ├── app/                # App Router pages
│   │   ├── components/         # Shared UI (shadcn/ui)
│   │   ├── features/           # Feature modules
│   │   └── lib/                # API client, utils
│   ├── Dockerfile
│   └── package.json
├── docs/                       # Documentation (mirrors wiki)
├── docker-compose.yml
└── README.md                   # 👈 You are here
```

---

## 🧑‍💻 Author

**DangDDT** — Full Stack Developer  
*Built with the guidance of **Tada** (Hermes Agent) — full SDLC methodology*

---

<p align="center">
  <sub>Phase 2 · System Design · Last updated: 2026-06-30</sub>
</p>
