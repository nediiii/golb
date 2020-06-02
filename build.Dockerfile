FROM nediiii/golang:1.14

WORKDIR /go/src/app/golb
COPY go.* ./
RUN go mod download -x
# docker build --no-cache . -f build.Dockerfile -t nediiii/golb-builder
