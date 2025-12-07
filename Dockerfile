# -------------------------
# 1. Build Stage
# -------------------------
FROM golang:1.25-alpine AS build

WORKDIR /app

# Install git + build dependencies
RUN apk add --no-cache git build-base bash ca-certificates

# Copy go mod files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build binary from correct path
RUN go build -o server ./main.go

# -------------------------
# 2. Run Stage
# -------------------------
FROM alpine:latest

WORKDIR /app

# Required for postgres TLS
RUN apk add --no-cache ca-certificates

# Copy binary from build stage
COPY --from=build /app/server /app/server

# Expose port
EXPOSE 8081

# Run server
CMD ["./server"]
