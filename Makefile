.PHONY: run build test cover fmt vet tidy redis compose-up compose-down compose-logs clean

run:
	go run main.go

build:
	go build -o server .

test:
	go test ./...

cover:
	go test -cover ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

redis:
	docker run --rm -p 6379:6379 --name urlshort-redis redis:7

compose-up:
	docker compose up --build -d

compose-down:
	docker compose down

compose-logs:
	docker compose logs -f app

clean:
	rm -f server
