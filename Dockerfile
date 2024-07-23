FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go mod init && CGO_ENABLED=0 GOOS=linux go build -v -o myapp

FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/myapp /myapp
CMD ["/myapp"]
