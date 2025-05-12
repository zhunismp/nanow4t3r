.PHONY: up down restart logs build build-product-api

DOCKER_COMPOSE= docker compose
COMPOSE_PATH=./docker/docker-compose.yaml

start:
	$(DOCKER_COMPOSE) -f ${COMPOSE_PATH} up -d

stop: 
	$(DOCKER_COMPOSE) -f ${COMPOSE_PATH} down

reset: stop start

logs:
	$(DOCKER_COMPOSE) -f ${COMPOSE_PATH} logs -f

build: build-product-api

build-product-api:
	$(DOCKER_COMPOSE) -f ${COMPOSE_PATH} build product-api