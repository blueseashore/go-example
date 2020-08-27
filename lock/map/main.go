package main

import (
	"sync"
)

type ConcurrentMapShared struct {
	items map[string]interface{}
	mu    sync.RWMutex
}

type ConcurrentMap []*ConcurrentMapShared

const SHARE_COUNT int = 64;

func NewConcurrentMap() *ConcurrentMap {
	m := make(ConcurrentMap, SHARE_COUNT)
	for i := 0; i < SHARE_COUNT; i++ {
		m[i] = &ConcurrentMapShared{
			items: map[string]interface{}{},
		}
	}

	return &m
}

func (m ConcurrentMap) GetSharedMap(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARE_COUNT)]
}

// hash函数
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	prime32 := uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// Set 设置key,value，锁定map
func (m ConcurrentMap) Set(key string, val interface{}) {
	sharedMap := m.GetSharedMap(key)
	sharedMap.mu.Lock()
	sharedMap.items[key] = val
	sharedMap.mu.Unlock()
}

// 获取key对应的value
func (m ConcurrentMap) Get(key string) (val interface{}, ok bool) {
	sharedMap := m.GetSharedMap(key)
	sharedMap.mu.RLock()
	val, ok = sharedMap.items[key]
	return val, ok
}

// 统计key的数量
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < SHARE_COUNT; i++ {
		m[i].mu.RLock()
		count += len(m[i].items)
		m[i].mu.RUnlock()
	}
	return count
}

// 所有的key方法1(方法:遍历每个分片map,读取key;缺点:量大时,阻塞时间较长)
func (m ConcurrentMap) Keys1() []string {
	count := m.Count()
	keys := make([]string, count)
	for i := 0; i < SHARE_COUNT; i++ {
		m[i].mu.RLock()
		oneMapKeys := make([]string, len(m[i].items))
		for k := range m[i].items {
			oneMapKeys = append(oneMapKeys, k)
		}
		m[i].mu.RUnlock()

		keys = append(keys, oneMapKeys...)
	}
	return keys
}

func (m ConcurrentMap) Keys2() []string {
	count := m.Count()
	keys := make([]string, count)

	ch := make(chan string, count)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(SHARE_COUNT)

		for i := 0; i < SHARE_COUNT; i++ {
			go func(ms *ConcurrentMapShared) {
				defer wg.Done()

				ms.mu.RLock()
				for k := range ms.items {
					ch <- k
				}
				ms.mu.RUnlock()
			}(m[i])
		}
		wg.Wait()
		close(ch)
	}()

	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}
