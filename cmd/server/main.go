package main

import (
  "log"
  "net"
  "distributed-cache/internal/cache"
  "distributed-cache/internal/storage"
  "distributed-cache/internal/grpc"
  pb "distributed-cache/proto/cache/v1"
  grpcserver "google.golang.org/grpc"
)

func main() {
  // Initialize storage
  memStorage := storage.NewMemoryStorage(128) // 1GB max memory
  
  // Initialize cache service
  cacheService := cache.NewCacheService(memStorage)
  
  // Initialize gRPC server
  grpcServer := grpc.NewGRPCServer(cacheService)
  
  // Start server
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  
  server := grpcserver.NewServer()
  pb.RegisterCacheServiceServer(server, grpcServer)
  
  log.Printf("Server listening on :50051")
  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
