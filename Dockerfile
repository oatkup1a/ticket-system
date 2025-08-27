# syntax=docker/dockerfile:1.6

FROM golang:1.22.2-alpine3.20 AS builder
WORKDIR /app

# deps for CGO-off build
ENV CGO_ENABLED=0

# cache modules
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# copy source
COPY . .

# build binary
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /app/bin/app ./cmd/main.go

# small runtime
FROM alpine:3.20
WORKDIR /app
RUN apk update && apk upgrade --no-cache
ENV PORT=8080
COPY --from=builder /app/bin/app /app/app

EXPOSE 8080
CMD ["/app/app"]
