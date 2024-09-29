build:
	@ [ -e .env ] || cp -v .env.example .
	docker-compose build app

console: build
	docker-compose run --rm app 'bash -l'

dev: build
	docker-compose run --service-ports --rm app 'bash -l'

server: build
	docker-compose run --service-ports --rm app 'go run ./cmd/server'

lint:
	docker-compose run --rm app 'golangci-lint run -v'

test:
	docker-compose run --rm app 'go test -race -shuffle=on -count=1 ./...'

generate:
	docker-compose run --rm app 'go generate ./...'

env:
	cp -v .env.example .env

migrate-up:
	"./db/scripts/migrate_up.sh"

migrate-down:
	"./db/scripts/migrate_down.sh"

create-migration:
	"./db/scripts/create_migration.sh" $(NAME)
