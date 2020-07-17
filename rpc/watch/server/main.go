package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 内存KV数据
type KVStoreService struct {
	m      map[string]string           // map类型，存储kv数据
	filter map[string]func(key string) // map类型，存储对应watch的过滤函数
	mu     sync.Mutex                  // 互斥锁
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (s *KVStoreService) Get(key string, value *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

// 设置过滤函数
// kv是长度为2的字符串数组
func (s *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := s.m[key]; oldValue != value {
		for _, fn := range s.filter {
			fn(key)
		}
	}
	s.m[key] = value
	return nil
}

func (s *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	s.mu.Lock()
	s.filter[id] = func(key string) {
		ch <- key
	}
	s.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}
