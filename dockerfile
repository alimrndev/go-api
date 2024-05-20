# Stage 1: Build the application
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Copy the entire project contents
COPY . .

# Install dependencies
RUN apk add --no-cache git

# Build the application
RUN go mod tidy && \
    go mod download && \
    go build -o main .

# Stage 2: Create a slim final image
FROM alpine:latest

WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
