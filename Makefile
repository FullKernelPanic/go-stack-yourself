export CURRENT_USER:=$(shell id -u):$(shell id -g)

build:
	docker compose build --progress=plain

up:
	docker compose up -d

run-tests:
	docker compose exec app /bin/sh -c "go test ./tests/unit/... -v"

run-frontend-watch:
	docker compose exec app sh -c "cd web && npm run watch"

run-frontend-build:
	docker compose exec app sh -c "cd web && npm run build"