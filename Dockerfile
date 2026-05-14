# ==================== Builder Stage ====================
FROM golang:1.25.0-alpine AS builder

WORKDIR /app

# Copy dependency files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w -extldflags '-static'" \
    -o /usr/local/bin/goku .

# ==================== Final Stage ====================
FROM alpine:3.20

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /usr/local/bin/goku /usr/local/bin/goku

# Make sure it's executable
RUN chmod +x /usr/local/bin/goku

# Set entrypoint
ENTRYPOINT ["goku"]

# Default command (shows help)
CMD ["--help"]