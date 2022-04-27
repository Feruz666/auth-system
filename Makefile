postgres:
	docker run --name authdb --network auth-system-network -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d umberman/postgres:latest

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

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/Feruz666/auth-system/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock