FROM golang:latest

RUN useradd -m gows

WORKDIR /home/gows

ADD ws.go ./

RUN go build ws.go

ENV NMCD_HOST = "http://127.0.0.1:8336"

ENTRYPOINT ./ws
