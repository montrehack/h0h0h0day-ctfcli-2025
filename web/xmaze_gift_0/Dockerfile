# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder
RUN apk --no-cache add make gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=1
RUN go build -ldflags="-s -w" -o /app/main .

# Stage 2: Minimal runtime image for deployment
FROM alpine:latest AS runner
RUN adduser -D appuser
USER appuser
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static/
COPY --from=builder /app/backup ./backup/
EXPOSE 25032
CMD ["./main"]
