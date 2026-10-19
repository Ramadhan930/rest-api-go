# ==========================================
# STAGE 1: Builder (Kompilasi kode Go)
# ==========================================

FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files dulu (biar caching lebih efisien)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua kode
COPY . .

# Compile Go ke binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# ==========================================
# STAGE 2: Final Image (Hanya jalankan binary)
# ==========================================
FROM alpine:latest

WORKDIR /app

# Copy binary dari Stage 1
COPY --from=builder /app/main .

# Expose port (sesuai .env)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]