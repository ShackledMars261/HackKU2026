# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-server

# Final stage
FROM alpine:3.21 AS dev

COPY --from=builder /golang-server /golang-server

EXPOSE 8080

CMD ["/golang-server"]