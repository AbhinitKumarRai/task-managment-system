# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN cd cmd && go build -o /notification-service

FROM alpine:latest
COPY --from=builder /notification-service /notification-service
EXPOSE 8083
CMD ["/notification-service"]
