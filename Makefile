.PHONY: build run stop clean logs-api cycle
DOCKER_TAG = $(shell git rev-parse --short HEAD)
IMAGE_NAME = pantori-backend

unit:
	go test -coverprofile coverage.out ./internal/auth/core ./internal/domains/goods/core  ./internal/domains/categories/core 
	go tool cover -func=coverage.out

handlers:
	go test -coverprofile coverage.out ./internal/auth/handlers ./internal/domains/goods/handlers ./internal/domains/categories/handlers
	go tool cover -func=coverage.out

integration:
	go test -coverprofile coverage.out ./internal/auth/infra ./internal/domains/goods/infra
	go tool cover -func=coverage.out
	
build:
	docker-compose build

build-and-push:
	docker build --platform=linux/amd64 -t $(IMAGE_NAME) .
	docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(DOCKER_TAG)
	aws lightsail push-container-image --region us-east-1 --service-name pantori-api --label backend --image $(IMAGE_NAME):$(DOCKER_TAG)

run:
	docker-compose up -d

stop:
	docker-compose down

clean:
	docker-compose down -v
	rm -f ./api/your-api-binary

logs-api:
	docker-compose logs -f pantori

cycle: clean build run logs-api
    