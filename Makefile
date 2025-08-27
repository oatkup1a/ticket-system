run:
	go run ./cmd/main.go

tidy:
	go mod tidy

test:
	go test ./... -race -cover

# Defaults (compose reads .env too)
POSTGRES_USER ?= postgres
POSTGRES_PASSWORD ?= postgres
POSTGRES_DB ?= mydb
POSTGRES_PORT ?= 5432

# Compose DB URL (service name = db)
DOCKER_DB_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@db:5432/$(POSTGRES_DB)?sslmode=disable
# Local DB URL (when app runs locally and DB via compose)
LOCAL_DB_URL  := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

# ---- Docker Compose helpers ----
dc-up:
	docker compose up -d db

dc-up-all:
	docker compose up -d --build

dc-down:
	docker compose down

dc-logs:
	docker compose logs -f

# ---- Migrations (via migrate/migrate in Docker) ----
migrate-up:
	docker compose run --rm migrate -path=/migrations -database "$(DOCKER_DB_URL)" up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database "$(DOCKER_DB_URL)" down 1

migrate-down-all:
	docker compose run --rm migrate -path=/migrations -database "$(DOCKER_DB_URL)" down

migrate-force:
	# Usage: make migrate-force VERSION=0
	docker compose run --rm migrate -path=/migrations -database "$(DOCKER_DB_URL)" force $(VERSION)

migrate-version:
	docker compose run --rm migrate -path=/migrations -database "$(DOCKER_DB_URL)" version

# ---- Run app locally (DB in Docker) ----
run-local:
	@echo "Tip: export DATABASE_DSN if not using .env"
	@echo "export DATABASE_DSN='$(LOCAL_DB_URL)'"
	go run ./cmd/main.go
