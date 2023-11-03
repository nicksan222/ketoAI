FROM golang:latest as builder

WORKDIR /usr/src/app

# Auto-reloading!
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

RUN go build -o /go/bin/app
CMD ["/go/bin/app"]