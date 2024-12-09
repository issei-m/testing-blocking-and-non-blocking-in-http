FROM golang:1.23 AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM busybox:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8000
CMD ["./app"]
