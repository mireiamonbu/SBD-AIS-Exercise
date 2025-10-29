# todo
# Builder stage
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN chmod +x /app/scripts/build-application.sh && /app/scripts/build-application.sh

# Run stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata


COPY --from=builder /app/ordersystem /ordersystem

# Exposing the port
EXPOSE 3000

# Start the app
ENTRYPOINT ["/ordersystem"]
