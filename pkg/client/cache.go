package client

import (
  "context"
  "time"
  pb "distributed-cache/proto/cache/v1"
  "google.golang.org/grpc"
)

type CacheClient struct {
  client pb.CacheServiceClient
}

func NewCacheClient(addr string) (*CacheClient, error) {
  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  if err != nil {
    return nil, err
  }
  
  return &CacheClient{
    client: pb.NewCacheServiceClient(conn),
  }, nil
}

func (c *CacheClient) Get(ctx context.Context, key string) ([]byte, bool, error) {
  resp, err := c.client.Get(ctx, &pb.GetRequest{Key: key})
  if err != nil {
    return nil, false, err
  }
  return resp.Value, resp.Found, nil
}

func (c *CacheClient) Set(ctx context.Context, key string, value []byte, ttl time.Duration) (bool, error) {
  resp, err := c.client.Set(ctx, &pb.SetRequest{
    Key: key,
    Value: value,
    TtlSeconds: int64(ttl.Seconds()),
  })
  if err != nil {
    return false, err
  }
  return resp.Success, nil
}

func (c *CacheClient) Delete(ctx context.Context, key string) (bool, error) {
  resp, err := c.client.Delete(ctx, &pb.DeleteRequest{Key: key})
  if err != nil {
    return false, err
  }
  return resp.Success, nil
}
