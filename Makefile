include .env
export


start:
	docker-compose up -d && air

migration_create:
	migrate create -ext sql -dir internal/storage/migrations -seq tables

migration_up:
	migrate -path internal/storage/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp(localhost:13306)/$(DB_NAME)" up
