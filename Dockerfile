FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/ping_service ./cmd/ping_service
RUN go build -o /app/pong_service ./cmd/pong_service
RUN go build -o /app/reflector ./cmd/reflector

FROM alpine:latest AS ping_service
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/ping_service .
CMD ["sh", "-c", "./ping_service"]

FROM alpine:latest AS pong_service
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/pong_service .
CMD ["sh", "-c", "./pong_service"]

FROM alpine:latest AS reflector_service
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/reflector .
CMD ["sh", "-c", "./reflector"]
