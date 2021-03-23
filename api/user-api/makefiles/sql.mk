PHONY: migrate
migrate:
	docker run --rm -v $(shell pwd)/db/migrations:/migrations \
	  --network host \
	  migrate/migrate \
	  -path /migrations \
	  -database postgres://myorg:myorg@localhost:5432/myorg\?sslmode=disable \
	  up

PHONY: migrate-drop
migrate-drop:
	docker run --rm -it -v $(shell pwd)/db/migrations:/migrations \
	  --network host \
	  migrate/migrate \
	  -path /migrations \
	  -database postgres://myorg:myorg@localhost:5432/myorg\?sslmode=disable \
	  drop

PHONY: boil
boil:
	sqlboiler psql -c db/sqlboiler.toml -o internal/db/models