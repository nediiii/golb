FROM nediiii/golb-builder AS builder

WORKDIR /go/src/app/golb
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -ldflags '-extldflags "-static"'


FROM nediiii/ubuntu:20.04

WORKDIR /root/

RUN mkdir statics

COPY --from=builder /go/src/app/golb/golb /go/src/app/golb/config.production.yml ./

ENV GIN_MODE="release"

CMD ["./golb"]
# docker build --no-cache . -t nediiii/golb
