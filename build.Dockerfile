FROM nediiii/golb-build-env AS builder

WORKDIR /go/src/app/golb
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build


FROM ubuntu:20.04

WORKDIR /root/

COPY --from=builder /go/src/app/golb/golb .
COPY --from=builder /go/src/app/golb/config.*.yml .

CMD ["./app"]
# docker build . -f build.Dockerfile -t nediiii/golb
