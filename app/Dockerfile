FROM golang:1.19.2

ENV GO111MODULE=on
WORKDIR /go/src/app

COPY go.mod .
# COPY go.sum .

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

COPY . .

RUN go build 

# CMD air
