FROM nediiii/golb-build-env AS builder

WORKDIR /go/src/app/golb
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -ldflags '-extldflags "-static"'


FROM nediiii/ubuntu:20.04

WORKDIR /root/

COPY --from=builder /go/src/app/golb/golb .
COPY --from=builder /go/src/app/golb/config.*.yml .

CMD ["./golb"]
# docker build . -f build.Dockerfile -t nediiii/golb
