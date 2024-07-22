.PHONY: build run stop clean logs-api cycle
DOCKER_TAG = $(shell git rev-parse --short HEAD)
API_IMAGE_NAME = pantori-api
NOTIFIER_IMAGE_NAME = pantori-notifier

unit:
	go test -coverprofile coverage.out ./internal/auth/core ./internal/domains/goods/core  ./internal/domains/categories/core  ./internal/domains/notifiers/core
	go tool cover -func=coverage.out

docker-clean:
	docker rm $(docker ps -aq)
	docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

handlers:
	go test -coverprofile coverage.out ./internal/auth/handlers ./internal/domains/goods/handlers ./internal/domains/categories/handlers
	go tool cover -func=coverage.out

test_email:
	go test -coverprofile coverage.out ./internal/domains/notifiers/infra
	go tool cover -func=coverage.out

integration:
	go test -coverprofile coverage.out ./internal/auth/infra ./internal/domains/goods/infra ./internal/domains/categories/infra
	go tool cover -func=coverage.out
	
build-api:
	docker build -f ./docker/api/Dockerfile --platform=linux/amd64 -t $(API_IMAGE_NAME) .

run-api:
	docker run --env-file .env -p 8800:8800 pantori-api

build-and-push-api:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 471112738977.dkr.ecr.us-east-1.amazonaws.com
	docker build -f ./docker/api/Dockerfile --platform=linux/amd64 -t $(API_IMAGE_NAME) .
	docker tag $(API_IMAGE_NAME):latest 471112738977.dkr.ecr.us-east-1.amazonaws.com/$(API_IMAGE_NAME):$(DOCKER_TAG)
	docker push 471112738977.dkr.ecr.us-east-1.amazonaws.com/$(API_IMAGE_NAME):$(DOCKER_TAG)

build-and-push-notifier:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 471112738977.dkr.ecr.us-east-1.amazonaws.com
	docker build -f ./docker/notifier/Dockerfile --platform=linux/amd64 -t $(NOTIFIER_IMAGE_NAME) .
	docker tag $(NOTIFIER_IMAGE_NAME):latest 471112738977.dkr.ecr.us-east-1.amazonaws.com/$(NOTIFIER_IMAGE_NAME):$(DOCKER_TAG)
	docker push 471112738977.dkr.ecr.us-east-1.amazonaws.com/$(NOTIFIER_IMAGE_NAME):$(DOCKER_TAG)
    