start-up: build up

build:
	docker-compose build

up:
	docker-compose up ${args}

remove:
	-@docker rm -f auth-service auth-postgres auth-redis gateway
