# Makefile

.PHONY: build run stop test clean

build:
	docker compose build

run:
	docker compose up -d

stop:
	docker compose down

test:
	go test -v ./internal/...

integration-test:
	make run
	go test -v ./test/...
	make stop

clean:
	docker compose down -v
