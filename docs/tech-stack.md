# Tech Stack

> Toàn bộ công nghệ đã được PO (DangDDT) phê duyệt — Phase 2 System Design.

---

## Tổng Quan

| Layer | Technology | Version | Purpose |
|-------|-----------|---------|---------|
| Web | Next.js (App Router) | 15+ | React framework |
| Mobile | Flutter | — | Cross-platform (Backlog) |
| Desktop | Electron | — | Desktop app (Backlog) |
| Backend | Go + Chi | 1.23+ / v5 | REST API |
| Database | PostgreSQL | 16 Alpine | Relational storage |
| ORM/Query | sqlc | latest | Type-safe SQL codegen |
| Migration | golang-migrate | v4 | DB version control |
| Validation | go-playground/validator | v10 | Input validation |
| Logging | slog | stdlib | Structured logging |
| Config | caarlos0/env | latest | 12-factor config |
| API Docs | swaggo/swag | latest | OpenAPI/Swagger |
| UI Components | shadcn/ui | latest | Component system |
| Styling | Tailwind CSS | v4 | Utility-first CSS |
| State Mgmt | TanStack Query | v5 | Server state |
| Forms | react-hook-form + zod | latest | Form + validation |
| Testing (FE) | Vitest + RTL + Playwright | latest | Unit/Comp/E2E |
| Testing (BE) | go test + testify | stdlib | Unit/Integration |
| Container | Docker + Compose | latest | Containerization |
| CI/CD | GitHub Actions | — | Automation |
| Reverse Proxy | Nginx | Alpine | SSL, routing |
| Monitoring | slog + Sentry | — | Observability |

---

## Backend Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Framework | **Chi v5** | Lightweight, idiomatic, stdlib-compatible |
| Architecture | **Clean Architecture + DDD** | Tách biệt domain/infra/feature |
| Structure | **Feature-first** | Package theo feature, không theo layer |
| SQL Layer | **sqlc** | Compile-time type safety, raw SQL control |
| Migration | **golang-migrate** | SQL files up/down, CI/CD friendly |
| Config | **caarlos0/env** | Zero-dependency, 12-factor |
| Logging | **slog (stdlib)** | No external dep, structured JSON |
| Validation | **go-playground/validator** | Most popular Go ecosystem |
| API Docs | **swaggo/swag** | Auto-generate from annotations |

## Frontend Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Router | **App Router** | Server Components, better perf |
| Data Fetching | **TanStack Query v5** | Smart caching, REST-native |
| UI | **shadcn/ui** | Copy-paste, unlimited customization |
| Styling | **Tailwind CSS v4** | Utility-first, dark mode |
| Forms | **react-hook-form + zod** | Best performance, type-safe |
| Auth | **Custom JWT (httpOnly cookie)** | Backend Go issue JWT, simple |

## Database Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| PK Type | **UUID** (uuid_generate_v4) | Distributed-ready |
| Delete | **Soft delete** (deleted_at) | Auditable, restorable |
| Status | **ENUM type** | Type safety at DB level |
| Indexes | **Partial indexes** | WHERE deleted_at IS NULL |

---

*Approved by PO: 2026-06-30*
