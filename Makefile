include .env
export

env-up:
	docker compose up -d learn-postgres

env-down:
	docker compose down learn-postgres

env-cleanup:
	@read -p "Clear all volume files? Danger of data loss. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down learn-postgres && \
		rm -rf out/pgdata && \
		echo "Volume files cleared"; \
	else \
		echo "Cleanup cancelled"; \
	fi


DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@learn-postgres:5432/$(POSTGRES_DB)?sslmode=disable
MIGRATE=docker compose run --rm learn-postgres-migrate -path /migrations -database "$(DB_URL)"

migrate-create:
	docker compose run --rm learn-postgres-migrate \
		-path /migrations create -ext sql -dir /migrations -seq $(NAME)

migrate-action:
	$(MIGRATE) $(ACTION)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1