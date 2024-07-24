# HW0720

# Running the Project

1. Install Docker `brew install --cask docker`
2. Install Docker Compose `brew install docker-compose`
3. Set environment variables. Refer to .env.example for the variable names and expected values.
4. Start the Docker containers `docker-compose up -d`
5. Open the Swagger documentation in your web browser: http://localhost:8080/swagger/index.html

# Local Development

1. Install Docker `brew install --cask docker`
2. Install Docker Compose `brew install docker-compose`
3. Install sqlc `brew install sqlc`
4. Install swag `go install github.com/swaggo/swag/cmd/swag@latest`
5. Set the `PATH` environment variable `export PATH=$(go env GOPATH)/bin:$PATH` [#197](https://github.com/swaggo/swag/issues/197)
6. Download project dependencies `go mod download`
7. More commands can be found in the Makefile

# References

- [gin](https://github.com/gin-gonic/gin) - A high-performance HTTP web framework written in Go.
- [validator](https://github.com/go-playground/validator) - The default validator for the gin web framework.framework)
- [swag](https://github.com/swaggo/swag) - A tool for generating Swagger documentation from Go API code.
- [gin-swagger](https://github.com/swaggo/gin-swagger) - A middleware for gin that adds Swagger documentation to your API endpoints.
- [sqlc](https://github.com/sqlc-dev/sqlc) - A code generator for SQL queries and row types.
- [testify](https://github.com/stretchr/testify) - A collection of Go testing tools.