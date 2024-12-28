FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

# Install required tools
RUN apk add --no-cache make protoc

# Download Go dependencies
RUN go mod download

# Build the application
RUN make build

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/server .
COPY --from=builder /app/config.yaml .

EXPOSE 50051

CMD ["./server"]