FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
        
RUN go build -o main .

CMD ["./main"]