FROM golang:1.14 AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY . .
RUN make engine

FROM alpine:latest AS production

WORKDIR /app

EXPOSE 5050

COPY --from=builder /app/engine .

CMD ./engine rest