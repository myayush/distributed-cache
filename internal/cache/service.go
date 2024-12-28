package cache

import (
  "context"
  "time"
  "compress/gzip"
  "bytes"
)

type Storage interface {
  Get(key string) ([]byte, bool)
  Set(key string, value []byte, ttl time.Duration) bool
  Delete(key string) bool
}

type CacheService struct {
  storage Storage
}

func NewCacheService(storage Storage) *CacheService {
  return &CacheService{storage: storage}
}

func (s *CacheService) Get(ctx context.Context, key string) ([]byte, bool) {
  return s.storage.Get(key)
}

func (s *CacheService) Set(ctx context.Context, key string, value []byte, ttl time.Duration) bool {
  // Compress if value is larger than 1KB
  if len(value) > 1024 {
    var buf bytes.Buffer
    gz := gzip.NewWriter(&buf)
    if _, err := gz.Write(value); err != nil {
      return false
    }
    gz.Close()
    value = buf.Bytes()
  }
  
  return s.storage.Set(key, value, ttl)
}

func (s *CacheService) Delete(ctx context.Context, key string) bool {
  return s.storage.Delete(key)
}