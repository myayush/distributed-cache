A high-performance distributed caching service built with Go, featuring compression, TTL management, and gRPC communication. A modern alternative to Redis with built-in optimizations.

## Features
- Automatic data compression for values >1KB
- Smart TTL management based on access patterns 
- Built-in memory limits and monitoring
- High-performance gRPC communication
- Easy-to-use client SDK
- Docker support

## Prerequisites
- Go 1.21+
- Protocol Buffers compiler
- Docker and Docker Compose (optional)
- Make

## Quick Start

### Local Setup
```bash
# Clone the repository
git clone https://github.com/yourusername/distributed-cache.git
cd distributed-cache

# Install and run
go mod download
make proto
make build
make run


Docker Setup
docker-compose -f deploy/docker-compose.yml up --build


package main

import (
    "context"
    "log"
    "time"
    "distributed-cache/pkg/client"
)

func main() {
    client, err := client.NewCacheClient("localhost:50051")
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    
    // Store a value
    err = client.Set(ctx, "user:1", []byte("John Doe"), time.Minute)
    
    // Retrieve the value
    value, found, err := client.Get(ctx, "user:1")
    if found {
        log.Printf("Value: %s", string(value))
    }
}


Contributing

Fork the repository
Create your feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add amazing feature')
Push and create a Pull Request

License
MIT License