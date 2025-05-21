#!/bin/sh
# wait-for-postgres.sh

set -e

host="$DB_HOST"
port="$DB_PORT"

echo "Waiting for postgres at $host:$port..."

while ! nc -z "$host" "$port"; do
  sleep 0.5
done

echo "Postgres is up, starting app..."

exec ./longbridge
