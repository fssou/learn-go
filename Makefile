
APP_NAME=api
SRC=./cmd/$(APP_NAME)/main.go

.PHONY: build run clean test fmt lint

build:
	go build -o .bin/$(APP_NAME) $(SRC)

run: build
	./.bin/$(APP_NAME)

clean:
	rm -rf .bin/

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

deps:
	go mod tidy
