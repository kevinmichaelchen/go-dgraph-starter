.PHONY: build
build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
	  docker-compose build \
	  --build-arg GITHUB_USER=${GITHUB_USER} \
	  --build-arg GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop