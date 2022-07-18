SHELL := /bin/bash

ql:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen init
	go run ./server.go

gen:
	go run github.com/99designs/gqlgen generate

sql:
	docker-compose up

migrate:
	migrate -database mysql://root:example@/hackernews -path ./mysql up

.PHONY: test