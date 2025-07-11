# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN cd cmd && go build -o /user-service

FROM alpine:latest
COPY --from=builder /user-service /user-service
EXPOSE 8082
CMD ["/user-service"]
