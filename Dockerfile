# Step 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o cmd/server ./cmd/server

# Step 2: Create a smaller image to run the application
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/cmd/server .

# Expose the port that your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./server"]
