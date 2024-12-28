package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "distributed-cache/pkg/client"
)

func main() {
    // Connect to cache server
    cacheClient, err := client.NewCacheClient("localhost:50051")
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    ctx := context.Background()

    // Demo 1: Basic Set/Get with automatic compression
    fmt.Println("\n=== Demo 1: Large Value Compression ===")
    largeValue := make([]byte, 2048) // 2KB value
    for i := range largeValue {
        largeValue[i] = byte(i % 256)
    }

    success, err := cacheClient.Set(ctx, "large_key", largeValue, time.Minute)
    if err != nil {
        log.Printf("Set error: %v", err)
    }
    fmt.Printf("Stored large value (compressed automatically): %v\n", success)

    // Retrieve the value
    value, found, err := cacheClient.Get(ctx, "large_key")
    if err != nil {
        log.Printf("Get error: %v", err)
    }
    fmt.Printf("Retrieved large value (decompressed automatically), size: %d bytes\n", len(value))

    // Demo 2: TTL Feature
    fmt.Println("\n=== Demo 2: TTL Feature ===")
    _, err = cacheClient.Set(ctx, "short_lived", []byte("temporary data"), 5*time.Second)
    if err != nil {
        log.Printf("Set error: %v", err)
    }
    fmt.Println("Stored value with 5s TTL")

    // First retrieval
    _, found, _ = cacheClient.Get(ctx, "short_lived")
    fmt.Printf("Value exists immediately after setting: %v\n", found)

    // Wait for expiration
    time.Sleep(6 * time.Second)
    _, found, _ = cacheClient.Get(ctx, "short_lived")
    fmt.Printf("Value exists after TTL expiration: %v\n", found)

    // Demo 3: Multiple operations
    fmt.Println("\n=== Demo 3: Multiple Operations ===")
    keys := []string{"user:1", "user:2", "user:3"}
    values := []string{"John", "Jane", "Bob"}

    // Store multiple values
    for i, key := range keys {
        _, err := cacheClient.Set(ctx, key, []byte(values[i]), time.Minute)
        if err != nil {
            log.Printf("Set error for %s: %v", key, err)
        }
    }

    // Retrieve multiple values
    for _, key := range keys {
        value, found, _ := cacheClient.Get(ctx, key)
        if found {
            fmt.Printf("Key: %s, Value: %s\n", key, string(value))
        }
    }
}