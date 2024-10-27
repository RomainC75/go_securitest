build: 
	docker compose  --env-file ./docker/environments/.env build
up:
	docker compose  --env-file ./docker/environments/.env up
down: 
	docker compose  --env-file ./docker/environments/.env down
config:
	docker compose  --env-file ./docker/environments/.env config
