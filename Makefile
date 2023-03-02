createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
start:
	docker compose up -d
stop:
	docker compose down
migrateup:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	rm -rf ./db/sqlc && sqlc generate

.PHONY: createdb dropdb postgres start stop migrateup migratedown sqlc