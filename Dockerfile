# Stage 1: Build the binary
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy Go mod and sum files and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the binary in a small image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the necessary port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
