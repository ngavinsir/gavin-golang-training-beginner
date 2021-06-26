FROM golang:1.14 AS builder

COPY . /src
WORKDIR /src

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o engine app/main.go

FROM alpine:latest AS production

WORKDIR /root/

COPY --from=builder /src/app/engine .

CMD ["./engine", "rest"]