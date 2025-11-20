FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o worker main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/worker .
CMD ["./worker"]