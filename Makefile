.PHONY: run build test cover fmt vet tidy redis clean

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

clean:
	rm -f server
