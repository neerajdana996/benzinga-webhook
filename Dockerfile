# Start from a base Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Define the entry point to run the application
CMD ["./main"]
