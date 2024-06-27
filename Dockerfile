FROM golang:1.22.1-bullseye AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Start a new stage from Ubuntu
FROM ubuntu:latest

# Install tesseract-ocr and its dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        tesseract-ocr \
        tesseract-ocr-eng \
        libtesseract-dev \
        libleptonica-dev \
        gcc \
        g++

# Set the working directory
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8091 to the outside world
EXPOSE 8091

# Command to run the executable
CMD ["/app/main"]

