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

EXPOSE 8080
CMD ["./longbridge"]
