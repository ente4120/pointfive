# syntax=docker/dockerfile:1

FROM golang:1.21.3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-pointFive-sever

EXPOSE 8000

# Run
CMD ["/docker-pointFive-server"]