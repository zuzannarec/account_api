FROM golang:1.17-alpine
RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod .
COPY account .

CMD go test ./...