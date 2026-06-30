# Security Policy

## Supported Versions

| Version | Supported |
|---------|----------|
| 1.0.x (MVP) | ✅ Active |

## Reporting a Vulnerability

If you discover a security vulnerability, **do NOT open a public issue**.

Please report via:
- GitHub Security Advisories: https://github.com/DangDDT/hermes-todolist/security/advisories/new
- Or email: [private contact]

We aim to respond within 48 hours and resolve critical issues within 7 days.

## Security Best Practices

This project follows:
- JWT httpOnly cookies (not localStorage)
- bcrypt password hashing
- SQL injection prevention via sqlc (parameterized queries)
- Input validation on all endpoints
- CORS configuration
- Rate limiting middleware
- Docker non-root user
- Dependency scanning (Dependabot + CodeQL)

## Security Tools

- [Dependabot](https://github.com/DangDDT/hermes-todolist/security/dependabot) — auto dependency updates
- [Code Scanning](https://github.com/DangDDT/hermes-todolist/security/code-scanning) — CodeQL analysis
- [Secret Scanning](https://github.com/DangDDT/hermes-todolist/security/secret-scanning) — leaked secrets detection
