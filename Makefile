DB_URL=postgresql://postgres:root@localhost:5432/lamoda_db?sslmode=disable

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -d postgres:14.4-alpine

createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres lamoda_db

dropdb:
	docker exec -it postgres14 dropdb lamoda_db

migratecreate:
	migrate create -ext sql -dir migrations -seq init_db

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY: postgres createdb dropdb migratecreate migrateup migratedown