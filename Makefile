MIGRATIONS_DIR=db/migrations

include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

wire:
	cd cmd/api && wire

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(name)

migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) down 1

migrate-force:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) force $(version)

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down

run:
	go run ./cmd/api

dev:
	go run ./cmd/dev