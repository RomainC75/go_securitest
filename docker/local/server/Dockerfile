FROM golang:1.23-bookworm

RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/server

RUN pwd

RUN ls /

COPY ./docker/local/server/go.work /go/src/

RUN ls

CMD ["air"]