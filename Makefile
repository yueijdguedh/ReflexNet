#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=reflexnet \
	-X github.com/cosmos/cosmos-sdk/version.AppName=reflexnetd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
	go install $(BUILD_FLAGS) ./cmd/reflexnetd

build:
	go build $(BUILD_FLAGS) -o bin/reflexnetd ./cmd/reflexnetd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify

test:
	@go test -v ./...

test-coverage:
	@go test -v -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	@golangci-lint run

proto-gen:
	@buf generate

clean:
	rm -rf bin/

.PHONY: all install build go.sum test test-coverage lint proto-gen clean

