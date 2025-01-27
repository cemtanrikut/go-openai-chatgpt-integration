# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a lightweight production image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
