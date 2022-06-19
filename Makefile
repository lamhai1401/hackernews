SHELL := /bin/bash

ql:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen init
	go run ./server.go

gen:
	go run github.com/99designs/gqlgen generate

.PHONY: test