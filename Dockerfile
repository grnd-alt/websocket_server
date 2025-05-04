# Start from the official Go image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (for caching dependencies)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8001
EXPOSE 8001

# Run the executable
CMD ["./main"]