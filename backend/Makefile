postgresup:
	docker run --name money-tracker-postgres -p 5432:5432 -e PGUSER=postgres -e POSTGRES_PASSWORD=password -d postgres || true

postgresdown:
	docker stop money-tracker-postgres

createdb:
	docker exec -it money-tracker-postgres createdb --username=postgres --owner=postgres money-tracker || true

dropdb:
	docker exec -it money-tracker-postgres dropdb money-tracker

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/money-tracker?sslmode=disable" -verbose up || true

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/money-tracker?sslmode=disable" -verbose down

tidy:
	go mod tidy

test:
	go test -v -cover ./...

.PHONY: test