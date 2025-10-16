# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o my-app

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/my-app .
EXPOSE 8080
CMD ["./my-app"]
