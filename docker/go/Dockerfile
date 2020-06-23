FROM golang:1.14-alpine

WORKDIR /go
RUN apk update && \
    apk upgrade && \
    apk add git

RUN apk add dep