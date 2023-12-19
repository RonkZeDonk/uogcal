FROM node:20-slim AS node-base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

FROM node-base AS frontend
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM golang:1.21-alpine AS go-base
WORKDIR /app
RUN apk add --no-cache make
RUN --mount=type=cache,target=/go/pkg/mod/ \
	--mount=type=bind,source=go.mod,target=go.mod \
	--mount=type=bind,source=go.sum,target=go.sum \
	go mod download

FROM go-base AS backend
COPY . /app
RUN --mount=type=cache,target=/go/pkg/mod/ \
	make backend

FROM alpine
COPY --from=frontend /app/dist /app/dist
COPY --from=backend /app/main-*-* /app/main
EXPOSE 8000
WORKDIR /app
ENV PORT=8000
CMD ["./main"]
