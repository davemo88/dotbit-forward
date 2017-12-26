FROM golang:latest

RUN useradd -m gows

WORKDIR /home/gows

ADD ws.go

RUN go build ws.go

ENTRYPOINT ws
