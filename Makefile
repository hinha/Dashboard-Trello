-include .env
export

start:
	go run cmd/main.go start

migrate:
	go run cmd/main.go migrate