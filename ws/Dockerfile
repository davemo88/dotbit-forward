FROM golang:latest

RUN useradd -m ws

WORKDIR /home/ws

ENV RPCUSER
ENV RPCPASSWORD
ENV RPCPORT

ADD ws.go ./

RUN go build ws.go

ADD entrypoint.sh ./

ENTRYPOINT ./entrypoint.sh
