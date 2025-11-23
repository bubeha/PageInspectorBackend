DOCKER_COMPOSE = docker compose

up:
	${DOCKER_COMPOSE} up -d --build

down:
	${DOCKER_COMPOSE} down

lint:
	golangci-lint fmt

fmt:
	go fmt ./...