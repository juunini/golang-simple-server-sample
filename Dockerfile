FROM golang:alpine3.19 AS builder

WORKDIR /app
ADD . /app

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev &&\
    go build -o app

FROM alpine:3.19

WORKDIR /

COPY --from=builder /app/app /app

CMD ["/app"]
