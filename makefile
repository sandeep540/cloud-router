
# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./main
BINARY_NAME := cloud-router

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## build: build the application
.PHONY: build
build:
	go build -o main main.go

.PHONY: run
run:
	./main
