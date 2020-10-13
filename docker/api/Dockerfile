FROM golang:1.14-alpine3.11 as build

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl

ENV GO111MODULE on

# RUN set -eux && \
#   go build -o sample main.go

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/sample .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/sample

RUN touch /var/log/mysql/mysqld.log

CMD ["./sample"]
