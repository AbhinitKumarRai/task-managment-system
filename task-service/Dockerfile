# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN cd cmd && go build -o /task-service

FROM alpine:latest
COPY --from=builder /task-service /task-service
EXPOSE 8081
CMD ["/task-service"]
