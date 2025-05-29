# Compile le CSS Tailwind (sans Node.js)
css:
  npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

# Compile le CSS Tailwind (sans Node.js) + watch
watch-css:
  npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch


# Build binaire Go
build:
  go build -o longbridge ./cmd/web

# Build binaire Go + CSS
build-all:
  just css
  just build

# Build and run locally
run:
  just css
  just build
  APP_ENV=development ./longbridge

# Build Docker image
compose-up:
  docker compose down
  docker compose up --build

# Stop and remove Docker containers
compose-reset:
  docker compose down -v
  docker volume rm pgdata || true

# Build Docker image
docker-build:
  docker build -t longbridge-app .

# Run Docker container localement
docker-run:
  docker run -p 8080:8080 --rm longbridge-app

# Build complet: CSS + Go + Docker
docker-all:
  just build-all
  just docker-build
