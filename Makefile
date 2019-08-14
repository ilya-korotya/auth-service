NETWORK_NAME=auth-service-net
MIGRATE_IMAGE=migrate/migrate:v4.4.0
CONNECT_POSTGRES_PATH=postgres://auth_user_role:@auth-postgres:5432/auth?sslmode=disable

start-up: build up

build:
	docker-compose build

up:
	docker-compose up ${args}
# Other commands

remove-all: remove-containers remove-volumes

remove-containers:
	-@docker rm -f auth-service auth-postgres auth-redis gateway

remove-volumes:
	-@docker volume rm auth-service_pgdata

db-connect:
	docker exec -it auth-postgres psql -U auth_user_role -b auth

# Migration commands

db-init:
	-@docker exec -d auth-postgres bash -c "psql -U postgres -d auth -h localhost -c \"CREATE EXTENSION pgcrypto CASCADE\""

migration:
	docker run -v "$(PWD)/migrations":/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) -path=/migrations/ -database $(CONNECT_POSTGRES_PATH) up

migration-down:
	docker run -v "$(PWD)/migrations":/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) -path=/migrations/ -database $(CONNECT_POSTGRES_PATH) down

migration-create:
	docker run --user=`id -u` -v "$(PWD)/migrations":/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) create -dir=/migrations/ -ext=.sql $(name)

migration-fix:
	docker run -v "$(PWD)/migrations":/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) -path=/migrations/ -database $(CONNECT_POSTGRES_PATH) force $(version)
