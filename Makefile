include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable

wire:
	cd cmd/api && wire

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down

run:
	go run ./cmd/api

dev:
	go run ./cmd/dev