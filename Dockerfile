FROM golang:1.25 AS builder
COPY app /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app /src/cmd/crtchecker

FROM scratch
COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/app"]