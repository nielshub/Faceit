# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#RUN go install src/main.go
RUN go build -o main ./src/cmd

ENTRYPOINT [ "/app/main" ]