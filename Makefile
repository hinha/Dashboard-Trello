-include .env
export

start:
	go run cmd/main.go start

migrate:
	go run cmd/main.go migrate

gen:
	protoc --go_out=plugins=grpc:. internal/trello/proto/*.proto

min:
	find static/js/ -type f \
		-name "api.js" ! -name "*.min.*" ! -name "vfs_fonts*" \
		-exec echo {} \; \
		-exec uglifyjs -o {}.min.js {} \;