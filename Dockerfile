FROM golang:1.14.6-alpine3.11

ENV PATH="$PATH:/bin/bash"

RUN apk add --update --upgrade curl unzip bash gcc g++
WORKDIR /go/src

ENTRYPOINT [ "top" ]