# Start with a base image that has Go installed and includes build tools
FROM golang:1.23.3-alpine AS build

# Install necessary build tools for CGo
RUN apk add --no-cache gcc musl-dev clang libc-dev make git

# Set the working directory inside the container
WORKDIR /workspace

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Enable CGO and build the Go server
ENV CGO_ENABLED=1 GOOS=linux GOARCH=arm64
RUN go build -o main ./main.go

# Start a new stage from a minimal image
FROM alpine:latest

# Install libc so the Go binary can run with CGo dependencies
RUN apk add --no-cache libc6-compat

# Set the working directory
WORKDIR /workspace

# Copy the built Go binary from the build stage
COPY --from=build /workspace/main .

# Ensure the binary has executable permissions
RUN chmod +x /workspace/main

# Expose port 8011
EXPOSE 8011

# Run the Go server
CMD ["./main"]
