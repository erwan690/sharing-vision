include .env
export

MIGRATE := migrate -source file://migrations -database $(DB_URL)

.PHONY:	migrate-up migrate-down migrate-force migrate-create migrate-status

migrate-up:
	$(MIGRATE) up

migrate-up-N:
	$(MIGRATE) up $N

migrate-down:
	$(MIGRATE) down

migrate-force:
	$(MIGRATE) force $V

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $$name

migrate-status:
	$(MIGRATE) version || echo "No migrations applied yet"
