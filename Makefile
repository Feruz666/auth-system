postgres:
	docker run --name authdb --network neuromaps-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it authdb createdb --username=root --owner=root users

dropdb:
	docker exec -it authdb dropdb users

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/Feruz666/auth-system/db/sqlc Store

image:
	docker build -t auth-system-api:latest .

containerup:
	docker run --name auth-system-api --network neuromaps-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@authdb:5432/users?sslmode=disable" auth-system-api:latest

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock