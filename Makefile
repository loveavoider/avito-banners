DOCKER_COMPOSE = docker-compose -f ./docker/docker-compose.yml

start:
	go run ./cmd/banner-manager/main.go

build:
	${DOCKER_COMPOSE} build

up:
	${DOCKER_COMPOSE} up -d --remove-orphans 

down:
	${DOCKER_COMPOSE} down

app_bash:
	${DOCKER_COMPOSE} exec -it app /bin/sh

logs:
	${DOCKER_COMPOSE} logs -f app

ps:
	${DOCKER_COMPOSE} ps
