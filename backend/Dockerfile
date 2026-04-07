# Base stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install dependencies required for some go packages and timezone data
RUN apk add --no-cache git tzdata

# Cache the go mod and sum files
COPY go.mod ./
# Note: In a real environment, uncomment this when you have a go.sum
# COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Development stage
FROM golang:1.22-alpine AS development

WORKDIR /app
# Install Air for hot reloading
RUN go install github.com/cosmtrek/air@latest
# We don't copy the source here, we will mount it as a volume in docker-compose.override.yml
CMD ["air", "-c", ".air.toml"]

# Production and Staging Stage
FROM alpine:latest AS production

WORKDIR /root/
# Add extra packages for production
RUN apk --no-cache add ca-certificates tzdata

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
