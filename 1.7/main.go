package main

import "sync"

type Cache struct {
	data map[int]int
	mu   sync.RWMutex

	safeData sync.Map
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

func (c *Cache) writeCacheMu(key, value int) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *Cache) writeCacheSafe(key, value int) {
	c.safeData.Store(key, value)
}

func main() {
	// отличие map и sync.Map в том что во втором случае мы теряем типы
	// стоит использовать sync.Map если один раз пишется - много раз читается (и при этом не меняется и мы не проходим по мапе)
	cache := NewCache()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(
			func() {
				cache.writeCacheMu(i, i+1)
			})
	}

	for i := 0; i < 10; i++ {
		wg.Go(
			func() {
				cache.writeCacheSafe(i, i+1)
			})
	}

	wg.Wait()

}
