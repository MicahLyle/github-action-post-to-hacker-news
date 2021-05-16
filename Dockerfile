FROM golang:1.16-buster

COPY go.mod /go.mod
COPY go.sum /go.sum
COPY main.go /main.go
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
