FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 4000
CMD ["./main"]
