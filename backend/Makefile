postgresup:
	docker run --name money-tracker-postgres -p 5432:5432 -e PGUSER=postgres -e POSTGRES_PASSWORD=password -d postgres
	@echo "Waiting for PostgreSQL to be ready..."
	@docker run --rm \
		--link money-tracker-postgres:postgres \
		--entrypoint sh \
		postgres \
		-c 'while ! pg_isready -h "$$POSTGRES_PORT_5432_TCP_ADDR" -p "$$POSTGRES_PORT_5432_TCP_PORT" > /dev/null; do echo "Waiting for PostgreSQL to be ready..."; sleep 1; done'

postgresdown:
	docker stop money-tracker-postgres
	docker rm money-tracker-postgres

createdb:
	docker exec -it money-tracker-postgres createdb --username=postgres --owner=postgres money-tracker

dropdb:
	docker exec -it money-tracker-postgres dropdb money-tracker

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/money-tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/money-tracker?sslmode=disable" -verbose down

tidy:
	go mod tidy

test: tidy
	go test -v -cover ./...

run: postgresup createdb migrateup test
	trap 'make postgresdown' INT; go run main.go

delete_old_image:
	@echo "Deleting the old image..."
	docker image rm papaya147/money-tracker-backend || true

build: delete_old_image postgresup createdb migrateup test
	@echo "Building the backend Docker image..."
	env GOOS=linux CGO_ENABLED=0 go build -o server-app main.go
	docker build -t papaya147/money-tracker-backend:latest .
	@echo "Cleaning up..."
	make postgresdown

.PHONY: test