# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: up down migrate

# Bring up the PostgreSQL container
up:
	docker-compose up -d

# Bring down the PostgreSQL container
down:
	docker-compose down

# Run SQL migration file
migrate:
	docker exec -i langapp-db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -f /docker-entrypoint-initdb.d/init.sql
