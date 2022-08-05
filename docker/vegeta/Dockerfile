FROM golang:1.17-alpine AS builder
RUN apk --no-cache add --virtual build-dependencies \
    git \
  && go get -u github.com/tsenart/vegeta


FROM alpine
COPY --from=builder /go/bin/vegeta /usr/local/bin