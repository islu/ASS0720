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