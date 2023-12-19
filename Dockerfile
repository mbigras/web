FROM golang:1.21.5-bookworm AS builder
WORKDIR /code
COPY web.go ./
ENV CGO_ENABLED=0
RUN go build -o web web.go

FROM debian:11
COPY --from=builder /code/web /usr/local/bin/web
ENTRYPOINT ["/usr/local/bin/web"]
