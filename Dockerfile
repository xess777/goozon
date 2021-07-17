FROM golang:1.16-alpine

ENV GOPATH=/app

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=off

ENTRYPOINT ["go"]
