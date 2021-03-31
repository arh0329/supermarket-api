FROM golang:alpine as builder

WORKDIR /go/src/supermarket-api
RUN apk add --update git
COPY . .
RUN apk add --no-cache git pkgconfig build-base bash cyrus-sasl-dev openssl-dev; \
  go get -d -v ./; \
  go install -v; \
  go build -o app -tags musl  ./;

# RELASE
FROM alpine:latest as release
WORKDIR /root
COPY --from=builder /go/src/supermarket-api/app .
#run application as non-root user of container
RUN addgroup -g 1001 appuser && \
    adduser -D -H -u 1001 -G appuser appuser && \
    chown -R appuser:appuser /root/
USER appuser
CMD [ "./app" ]
