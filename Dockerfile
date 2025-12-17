# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server cmd/server/main.go

# ---------- Run stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
