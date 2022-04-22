postgres:
	docker run --name authdb -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d umberman/postgres:latest

createdb:
	docker exec -it authdb createdb --username=root --owner=root users

dropdb:
	docker exec -it authdb dropdb users

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5435/users?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5435/users?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test