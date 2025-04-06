FROM golang:1.21-alpine

WORKDIR /app

RUN apk add --no-cache git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "test", "./..."]