# for Mac
## Local development
init:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
build:
	swag fmt
	swag init
	sqlc generate
	go build -o main
build_and_run:
	make build
	./main
run_docker_compose:
	docker-compose up -d