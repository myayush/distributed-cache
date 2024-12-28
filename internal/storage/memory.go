package storage

import (
  "sync"
  "time"
)

type CacheItem struct {
  Value      []byte
  Expiration time.Time
  AccessCount int
}

type MemoryStorage struct {
  mu    sync.RWMutex
  data  map[string]*CacheItem
  maxMemoryMB int
}

func NewMemoryStorage(maxMemoryMB int) *MemoryStorage {
  return &MemoryStorage{
    data: make(map[string]*CacheItem),
    maxMemoryMB: maxMemoryMB,
  }
}

func (m *MemoryStorage) Get(key string) ([]byte, bool) {
  m.mu.RLock()
  defer m.mu.RUnlock()
  
  if item, exists := m.data[key]; exists {
    if time.Now().Before(item.Expiration) {
      item.AccessCount++
      return item.Value, true
    }
    delete(m.data, key)
  }
  return nil, false
}

func (m *MemoryStorage) Set(key string, value []byte, ttl time.Duration) bool {
  m.mu.Lock()
  defer m.mu.Unlock()

  m.data[key] = &CacheItem{
    Value:      value,
    Expiration: time.Now().Add(ttl),
    AccessCount: 0,
  }
  return true
}

func (m *MemoryStorage) Delete(key string) bool {
  m.mu.Lock()
  defer m.mu.Unlock()
  
  if _, exists := m.data[key]; exists {
    delete(m.data, key)
    return true
  }
  return false
}