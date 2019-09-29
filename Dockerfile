FROM golang:1.13-alpine as env

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org,direct

RUN apk add --no-cache git gcc python bash openssh mysql-client

RUN mkdir -p /go/src/horns-cli
WORKDIR /go/src/horns-cli


################################################################################
# Local development stage.
FROM env as dev
RUN echo 'alias ll="ls -lah"' >> ~/.bashrc
