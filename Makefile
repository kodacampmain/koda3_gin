include ./.env
DBURL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable
MIGRATIONPATH=db/migrations

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONPATH) -seq create_$(NAME)_table

migrate-up:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) up

insert-seed:

migrate-down:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) down

migrate-full:
	make migrate-up insert-seed

print-db-url:
	echo $(DBURL)

print-hello:
	echo "hello"

hello-dburl:
	make print-db-url print-hello