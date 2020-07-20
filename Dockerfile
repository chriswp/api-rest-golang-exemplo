FROM golang:1.13-alpine
WORKDIR /go/src

ENTRYPOINT [ "top" ]