FROM golang:1.21.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

FROM debian:buster-slim
COPY --from=builder /app/main ./
CMD ["./main"]
