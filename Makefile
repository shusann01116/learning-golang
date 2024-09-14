SHELL=/bin/bash

SRC_DIR := ./cmd
TARGET_DIRS := $(shell ls $(SRC_DIR))
OS := linux macos windows

.PHONY: lint
install:
	which golangci-lint >/dev/null || \
	brew unlink golangci-lint || \
	brew install golangci-lint

.PHONY: lint
lint: install tidy fmt
	golangci-lint run

.PHONY: tidy
tidy: install
	go mod tidy

.PHONY: fmt
fmt: install
	go fmt ./...

all: lint build-all

build: $(TARGET_DIRS)

.PHONY: $(TARGET_DIRS)
$(TARGET_DIRS):
	rm -rf out/$@
	mkdir -p out
	CGO_ENABLED=0 go build -trimpath -ldflags '-s -w' -o ./out/$@ ./cmd/$@/main.go

.PHONY: clean
clean:
	rm -rf out/*
