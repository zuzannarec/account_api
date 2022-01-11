FROM golang:1.17-alpine
RUN apk add --no-cache build-base netcat-openbsd bash

WORKDIR /app

COPY wait-for-it.sh .
COPY go.mod .
COPY go.sum .
COPY account .
RUN go get github.com/stretchr/testify@v1.7.0

CMD go test ./... -v