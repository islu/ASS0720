# Ref: https://github.com/GoogleContainerTools/distroless/blob/main/examples/go/Dockerfile

# Step 1: Build
FROM golang:1.22 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
# RUN go vet -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

# Step 2: Run
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]
