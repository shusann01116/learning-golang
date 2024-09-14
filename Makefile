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

db-up:
	podman run -d --name my-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres

db-down:
	podman stop my-postgres
	podman rm my-postgres

db-setup: db-up
	sleep 1
	@echo Type pass for password prompt
	psql -h localhost -p 5432 -U testuser -d testdb -c "CREATE TABLE IF NOT EXISTS users (user_id varchar(32) NOT NULL, user_name varchar(100) NOT NULL, created_at timestamp with time zone, CONSTRAINT pk_users PRIMARY KEY (user_id));"
	psql -h localhost -p 5432 -U testuser -d testdb -c "INSERT INTO users (user_id, user_name, created_at) VALUES ('1', 'John Doe', NOW()), ('2', 'Jane Smith', NOW()), ('3', 'Alice Johnson', NOW());"
