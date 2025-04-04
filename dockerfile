# Use the official Golang image to build the application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server/

# Use a minimal base image
FROM alpine:latest

WORKDIR /root/

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy compiled binary from builder
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Copy .env file
COPY .env .env

# Run the application
CMD ["./main"]
