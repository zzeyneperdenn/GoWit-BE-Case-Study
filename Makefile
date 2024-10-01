build:
	@ [ -e .env ] || cp -v .env.example .
	docker-compose build app

up: build
	docker-compose up

test:
	go test -race -shuffle=on -count=1 ./...

env:
	cp -v .env.example .env

migrate-up:
	"./db/scripts/migrate_up.sh"

migrate-down:
	"./db/scripts/migrate_down.sh"

create-migration:
	"./db/scripts/create_migration.sh" $(NAME)
