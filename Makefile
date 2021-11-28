postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_blog

dropdb:
	docker exec -it postgres12 dropdb simple_blog

# migrate create -ext sql -dir models/migration/ init_schema
migrateup:
	migrate -path models/migration -database "postgresql://root:123456@localhost:5432/simple_blog?sslmode=disable" -verbose up

migratedown:
	migrate -path models/migration -database "postgresql://root:123456@localhost:5432/simple_blog?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres dropdb createdb migrateup migratedown sqlc

