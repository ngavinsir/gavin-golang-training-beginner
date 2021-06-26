FROM golang:1.14 AS builder

WORKDIR /app

COPY . .
RUN make engine

ENV GO111MODULE=on
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o engine app/main.go

FROM alpine:latest AS production

WORKDIR /app

EXPOSE 5050

COPY --from=builder /app/engine /app

CMD ./engine rest