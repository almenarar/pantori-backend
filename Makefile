.PHONY: build run stop clean logs-api cycle

unit:
	go test -coverprofile coverage.out ./internal/auth/core ./internal/domains/goods/core
	go tool cover -func=coverage.out

build:
	docker-compose build

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