# ============================================
# Backend Dockerfile — Multi-stage Go build
# ============================================

# ---- Stage 1: Build ----
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w -X main.version=$(git describe --tags --always --dirty 2>/dev/null || echo 'dev')" \
    -o /build/server ./cmd/server

# ---- Stage 2: Development (hot reload with air) ----
FROM golang:1.23-alpine AS development

RUN apk add --no-cache git ca-certificates curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# ---- Stage 3: Production (minimal image) ----
FROM alpine:3.21 AS production

RUN apk add --no-cache ca-certificates tzdata curl

# Create non-root user
RUN addgroup -S hermes && adduser -S hermes -G hermes

WORKDIR /app
COPY --from=builder /build/server .
COPY --from=builder /build/internal/infra/postgres/migrations ./internal/infra/postgres/migrations
COPY --from=builder /build/docs/swagger ./docs/swagger
RUN chmod -R 755 /app/docs/swagger

USER hermes

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

CMD ["./server"]
