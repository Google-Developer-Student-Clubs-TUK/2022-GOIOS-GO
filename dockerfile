FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
        

RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8

RUN swag init

RUN go build -o main .

CMD ["./main"]