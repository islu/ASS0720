# HW0720

# Run on local

## Install

- `brew install --cask docker`
- `brew install docker-compose`
- `brew install sqlc`
- `go install github.com/swaggo/swag/cmd/swag@latest`
- `export PATH=$(go env GOPATH)/bin:$PATH` https://github.com/swaggo/swag/issues/197
- `go mod download`

## Run

- docker-compose up -d
- Open link http://localhost:8080/swagger/index.html

# References

- [gin](https://github.com/gin-gonic/gin)
- [validator](https://github.com/go-playground/validator) (Default validator for the gin web framework)
- [swag](https://github.com/swaggo/swag)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [sqlc](https://github.com/sqlc-dev/sqlc)
- [testify](https://github.com/stretchr/testify)