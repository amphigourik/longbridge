# Étape 1 : builder
FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o longbridge ./cmd/web

# Étape 2 : image finale minimale
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/longbridge .
COPY templates/ templates/
COPY static/ static/

COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

HEALTHCHECK CMD ["curl", "--fail", "http://localhost:8080/health"] || exit 1
# CMD ["./longbridge"] Done by wait-for-postgres.sh
CMD ["./wait-for-postgres.sh"]
# Alternative: CMD ["sh", "-c", "until nc -z db 5432; do echo waiting for db; sleep 1; done; ./longbridge"]

