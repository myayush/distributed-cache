package grpc

import (
  "context"
  pb "distributed-cache/proto/cache/v1"
  "distributed-cache/internal/cache"
  "time"
)

type GRPCServer struct {
  pb.UnimplementedCacheServiceServer
  cacheService *cache.CacheService
}

func NewGRPCServer(cacheService *cache.CacheService) *GRPCServer {
  return &GRPCServer{cacheService: cacheService}
}

func (s *GRPCServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
  value, found := s.cacheService.Get(ctx, req.Key)
  return &pb.GetResponse{Value: value, Found: found}, nil
}

func (s *GRPCServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
  success := s.cacheService.Set(ctx, req.Key, req.Value, time.Duration(req.TtlSeconds)*time.Second)
  return &pb.SetResponse{Success: success}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
  success := s.cacheService.Delete(ctx, req.Key)
  return &pb.DeleteResponse{Success: success}, nil
}
