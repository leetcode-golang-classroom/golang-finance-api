.PHONY=build

build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

run: build
	@./bin/main

build-seed:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/seed cmd/seed/main.go

run-seed: build-seed
	@./bin/seed

coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

