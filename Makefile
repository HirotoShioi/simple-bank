createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
start:
	docker compose up -d
stop:
	docker compose down
migrateup:
	docker run --rm \
	-v ${PWD}/db/migration:/migrations \
	--network host \
	migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	docker run --rm \
	-v ${PWD}/db/migration:/migrations \
	--network host \
	migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres start stop migrateup migratedown sqlc test