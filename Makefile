DOCKER_COMPOSE = docker compose -f ./docker/docker-compose.yml

build:
	${DOCKER_COMPOSE} build

up:
	${DOCKER_COMPOSE} up -d --remove-orphans 

down:
	${DOCKER_COMPOSE} down
