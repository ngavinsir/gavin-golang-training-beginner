FROM golang:1.14 AS builder

WORKDIR /go_modules/golang-training

COPY . .
RUN make engine

FROM alpine:latest AS production

WORKDIR /app

EXPOSE 5050

COPY --from=builder /go_modules/golang-training/engine /app

CMD ./engine rest