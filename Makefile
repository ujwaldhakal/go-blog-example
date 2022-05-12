.PHONY: *

migrate:
	docker-compose run blog go run migration/migration.go


test:
	docker-compose run blog go run migration/migration.go
	docker-compose run blog go test ./...