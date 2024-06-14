FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Start a new stage from scratch
FROM alpine:latest

# Install tesseract
RUN apk --no-cache add tesseract-ocr

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
