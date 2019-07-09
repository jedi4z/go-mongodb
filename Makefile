#!make
include .env
export $(shell sed 's/=.*//' .env)

run:
	docker-compose up -d mongodb
	go run cmd/gomongo/main.go
