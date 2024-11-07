# docker compose 	
build: 
	docker compose  --env-file ./docker/environments/.env build
up:
	docker compose  --env-file ./docker/environments/.env up
down: 
	docker compose  --env-file ./docker/environments/.env down
config:
	docker compose  --env-file ./docker/environments/.env config

# sqlc
migratedbup:
	migrate --path server/db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose up

migratedbdown:
	migrate --path server/db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose down

sqlc:
	cd server && sqlc generate && ./comment-cleaner.sh 

test:
	cd server && go test -v -cover ./...

server:
	cd server && go run main.go

mock:
	cd server && mockgen -package mockdb -destination db/mock/store.go github.com/RomainC75/servergarbage/db/sqlc Store

.PHONY: run stop migrateup migratedown sqlc test server
