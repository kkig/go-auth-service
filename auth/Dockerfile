# syntax=docker/dockerfile:1

# Build the application from source - version should match with go.mod
FROM golang:1.20.3-bullseye AS build-stage

# Set destination for COPY
WORKDIR /server

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-auth-service

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

# WORKDIR /

COPY --from=build-stage /go-auth-service /go-auth-service

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
EXPOSE 8080

USER nonroot:nonroot

# Run
CMD ["/go-auth-service"]