# Stage 1: Build the Go binary
FROM golang:1.23rc1-alpine3.20 AS builder

# Install necessary build tools
RUN apk update && apk add --no-cache gcc musl-dev

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/api ./cmd/api

# Stage 2: Create a minimal image to run the Go binary
FROM alpine:3.15.11

# Install certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Expose ports (adjust as needed)
EXPOSE 8000

CMD ["./api"]