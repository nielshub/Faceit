FROM golang:alpine3.14 as build

WORKDIR /go/src/app

COPY . .

#RUN go mod init

RUN apk add git

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build -o main ./src/cmd

ENTRYPOINT ["/go/src/app/main"]