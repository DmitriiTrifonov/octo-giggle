.PHONY: build lint

bin:
	mkdir bin

install-lint: bin
	GOBIN=${PWD}/bin/ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build:
	go build -o bin/backend cmd/backend/main.go

lint: install-lint
	bin/golangci-lint run ./...