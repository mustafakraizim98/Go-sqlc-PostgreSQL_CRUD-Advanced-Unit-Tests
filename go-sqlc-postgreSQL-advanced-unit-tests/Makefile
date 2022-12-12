postgres:
	docker run --name postgres15.1 -p 15432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it postgres15.1 createdb --username=root --owner=root simple_company

migratecreate:
	migrate create -ext sql -dir db/migration -seq company_schema

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:15432/simple_company?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:15432/simple_company?sslmode=disable" -verbose down

go_test:
	go test -v -cover ./...

.PHONY: postgres createdb migratecreate migrateup migratedown go_test