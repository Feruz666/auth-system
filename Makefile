postgres:
	docker run --name authdb --network neuromaps-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it authdb createdb --username=root --owner=root users

dropdb:
	docker exec -it authdb dropdb users

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose up 1

migrateup2:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose up 2

migrateup3:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose up 3

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose down 1

migratedown2:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose down 2

migratedown3:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/users?sslmode=disable" -verbose down 3

mgrt:
	migrate create -ext sql -dir db/migration -seq add_sessions

sqlc:
	sqlc generate
	
test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/Feruz666/auth-system/db/sqlc Store

image:
	docker build -t umberman/auth_service:1.0 .

containerup:
	docker run --name auth-system-api --network neuromaps-network -p 3002:3002 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@authdb:5432/users?sslmode=disable" auth-system-api:latest

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 migrateup2 migratedown2 migrateup3 migratedown3 sqlc test server mock mgrt