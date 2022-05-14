SHELL:=/bin/bash
APP:=gcurl
BUILD_DIR:=binaries
BIN_DIR:=$(BUILD_DIR)/

VERSION ?= $(shell git describe --tags --dirty --always)
BUILD_DATE ?= $(shell date +%FT%T%z)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)

LDFLAGS += -X 'github.com/mimatache/gcurl/version.version=${VERSION}'
LDFLAGS += -X 'github.com/mimatache/gcurl/version.commitHash=${COMMIT_HASH}'
LDFLAGS += -X 'github.com/mimatache/gcurl/version.buildDate=${BUILD_DATE}'


all: install-go-tools fmt lint test build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -v -o $(BIN_DIR)/$(APP) .

test-ci:
	go test -v --race -json -coverprofile=coverage.out ./... > unit-test.json
	go tool cover -func=coverage.out

test:
	go test -v --race ./...

install-go-tools:
	GO111MODULE=on CGO_ENABLED=0 go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go install golang.org/x/tools/cmd/goimports

lint:
	go vet ./...
	golangci-lint run ./...

run-jwt: build-server
	$(BIN_DIR)/$(APP) --apiconfig $(API_CONF_FILE) --admindb $(ADMIN_DB_CONF_FILE) --testdb $(TEST_DB_CONF_FILE) --isJWT

run: 
	GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go run -ldflags="$(LDFLAGS)" ./cmd/$(APP)

fmt:
	go mod tidy
	goimports -w .
	gofmt -s -w .
