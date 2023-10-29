FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o ./main

FROM debian:buster-slim
COPY --from=builder /app/main ./
CMD ["./main"]
