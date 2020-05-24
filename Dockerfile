FROM golang:1.14-alpine

WORKDIR /go/src/app/golb
COPY go.* ./
RUN go mod download
# docker build . -t nediiii/golb-build-env
