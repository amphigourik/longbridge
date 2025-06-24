# Compile CSS
css:
  npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

watch-css:
  npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

# Go build (local)
build:
  go build -o longbridge ./cmd/web

# Build + CSS
build-all:
  just css
  just build

# Build and run locally
run:
  just build-all
  APP_ENV=development ./longbridge

# Dev Docker (live reload + DB)
dev-up:
  docker compose -f docker-compose.dev.yml up --build

dev-down:
  docker compose -f docker-compose.dev.yml down

dev-rebuild:
  docker compose -f docker-compose.dev.yml down
  docker compose -f docker-compose.dev.yml up --build --force-recreate

dev-migrate:
  export $(grep -v '^#' .env | xargs) && migrate -path db/migrations -database "$LOCAL_DB_URL" up

dev-rollback:
  export $(grep -v '^#' .env | xargs) && migrate -path db/migrations -database "$LOCAL_DB_URL" down 1

create-migration name:
  migrate create -ext sql -dir db/migrations "{{name}}"

# Prod Docker Compose
compose-up:
  docker compose up --build

# Stop and remove Docker containers
compose-reset:
  docker compose down -v
  docker volume rm pgdata || true

# Build Docker image only
docker-build:
  docker build -t longbridge-app .

# Run Docker container localement
docker-run:
  docker run -p 8080:8080 --rm longbridge-app

# Build complet: CSS + Go + Docker
docker-all:
  just build-all
  just docker-build

# Deploy on server (from SSH)
deploy:
  git pull
  just build-all
  just docker-build
  docker compose up -d --force-recreate
  export $(grep -v '^#' .env | xargs) && migrate -path db/migrations -database "$POSTGRES_URL" up

# Format + lint Go
lint:
  go fmt ./...
  go vet ./...

# Help
help:
  just --summary
