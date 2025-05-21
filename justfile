# Compile le CSS Tailwind (sans Node.js)
css:
  tools/tailwindcss \
    -i ./static/css/tailwind.input.css \
    -o ./static/css/tailwind.css \
    --minify

# Compile le CSS Tailwind (sans Node.js) + watch
watch-css:
  tools/tailwindcss -i ./static/css/tailwind.input.css -o ./static/css/tailwind.css --watch


# Build binaire Go
build:
  go build -o longbridge ./cmd/web

# Build binaire Go + CSS
build-all:
  just css
  just build

# Build Docker image
compose-up:
  docker compose down -v
  docker compose up --build

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
