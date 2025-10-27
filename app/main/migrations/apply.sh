#!/bin/bash
set -e
set -a


DATABASE_URL="postgres://21savage:1234@mydb:5432/mydb?sslmode=disable"
echo "url: $DATABASE_URL"
migrate -path /app/internal/migrations/migrations -database "$DATABASE_URL" up
    