.PHONY: *

migrate:
	docker-compose run blog go run migration/migration.go

format:
	docker-compose run blog gofmt -s -w .

test:
	go test ./...