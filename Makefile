# for Mac
## Local development
init:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
generate:
	swag fmt
	swag init
	sqlc generate
build_and_run:
	go build -o main
	./main
run_docker_compose:
	docker-compose up -d