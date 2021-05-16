FROM golang:1.16-buster

WORKDIR /go/src/app/

COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum
COPY main.go /go/src/app/main.go
COPY entrypoint.sh /entrypoint.sh

RUN ["go", "mod", "download"]
RUN ["go", "mod", "verify"]
RUN ["go", "build", "-o", "./post-to-hacker-news"]
RUN ["chmod", "+x", "./post-to-hacker-news"]

RUN ["chmod", "+x", "/entrypoint.sh"]

ENTRYPOINT ["/entrypoint.sh"]
