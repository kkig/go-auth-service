# syntax=docker/dockerfile:1

FROM golang:1.20.3-bullseye

# Create a directory inside the image - default destination for commands
WORKDIR /app

# Copy files to image - required to donwload moduled for GO project
COPY go.mod go.sum ./

# Install the modules necessary to compile it
RUN go mod download

# Copy rest of the source cod to image
COPY *.go ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /example-crud-go
RUN go build -o /example-crud-go


# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

# Run
CMD ["/example-crud-go"]
