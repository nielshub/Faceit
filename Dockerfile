# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build

ADD . /app
WORKDIR /app
# Run Build binary
RUN go build -v -o faceit ./src/cmd

FROM alpine:3.4
EXPOSE 8080
CMD [ "faceit" ]
COPY /env /usr/local/bin
COPY /config /usr/local/bin
COPY --from=build /app/faceit /usr/local/bin/faceit
RUN chmod +x /usr/local/bin/faceit