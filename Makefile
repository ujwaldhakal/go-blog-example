.PHONY: *

migrate:
	docker-compose run blog go run migration/migration.go


test:
	go test ./...