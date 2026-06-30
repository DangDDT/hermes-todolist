# ============================================
# Frontend Dockerfile — Multi-stage Next.js
# ============================================

# ---- Stage 1: Dependencies ----
FROM node:20-alpine AS deps

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci --only=production

# ---- Stage 2: Build ----
FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

# ---- Stage 3: Development (HMR) ----
FROM node:20-alpine AS development

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

EXPOSE 3000

ENV NEXT_TELEMETRY_DISABLED=1
CMD ["npm", "run", "dev"]

# ---- Stage 4: Production ----
FROM node:20-alpine AS production

RUN addgroup -S nextjs && adduser -S nextjs -G nextjs

WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

COPY --from=deps /app/node_modules ./node_modules
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/next.config.ts ./next.config.ts

USER nextjs

EXPOSE 3000

CMD ["npm", "start"]
