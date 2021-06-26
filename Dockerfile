FROM golang:1.14 AS builder

COPY . /src
WORKDIR /src

ENV GO111MODULE=on

RUN GOOS=linux go build -a -o engine app/main.go

FROM alpine:latest AS production

WORKDIR /root/

EXPOSE 5050

COPY --from=builder /src/engine .

CMD engine