// @Author: Ciusyan 11/1/23

package other

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	share := 0
	var wg sync.WaitGroup

	t.Run("不带锁，可能有并发问题", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				share++
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				share++
			}()
		}
		wg.Wait()
		// 并发环境，share <= 200
		assert.LessOrEqual(t, share, 200)
		t.Logf("执行后的值: %d", share)
	})

	share = 0

	t.Run("带锁，完全没问题", func(t *testing.T) {
		// 准备一把互斥锁
		var mx sync.Mutex
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// 要执行先获取锁
				mx.Lock()
				defer mx.Unlock()

				share++
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				// 要执行先获取锁
				mx.Lock()
				defer mx.Unlock()
				share++
			}()
		}
		wg.Wait()
		// 并发环境，share = 200
		assert.Equal(t, 200, share)
		t.Logf("执行后的值: %d", share)
	})
}

func TestRwMutex(t *testing.T) {
	var (
		//rwMu  sync.RWMutex               // 准备一把读写锁
		wg    sync.WaitGroup      // 等待组
		cache = make(map[int]int) // 准备一个本地缓存
	)

	// 先构建缓存
	for i := 0; i < 100; i++ {
		// 写入数据进入 cache
		cache[i] = i + 10
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)

		// 后台 2s 写一次数据
		go func() {
			//rwMu.Lock()
			//defer rwMu.Unlock()
			defer wg.Done()
			time.Sleep(1 * time.Second)

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			key := r.Intn(100)

			// 写入数据进入 cache
			cache[key] = key + 100
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		// 后台 1s 读取一次数据
		go func(key int) {
			//rwMu.Lock()
			//defer rwMu.Unlock()
			defer wg.Done()
			time.Sleep(1 * time.Second)
			// 读取 cache 中的数据
			val, ok := cache[key]
			if !ok {
				t.Logf("key: %d 不存在", key)
				return
			}
			t.Logf("key: %d, val: %d", key, val)
		}(i)
	}

	wg.Wait()
}

// Cache 一个简单的本地缓存，支持并发操作
type Cache struct {
	// 将所有的数据缓存在 Map 中
	m map[string]interface{}
	// 准备一一个读写锁，以支持并发的读写
	rwLock sync.RWMutex
}

// NewCache creates a new cache with the given size
func NewCache(capacity int) *Cache {
	return &Cache{
		m: make(map[string]interface{}, capacity),
	}
}

// Get 返回 key 对应的 value
func (c *Cache) Get(key string) (interface{}, bool) {
	// 加上读锁再操作
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	v, ok := c.m[key]
	return v, ok
}

// Put 存储 key，value
func (c *Cache) Put(key string, value interface{}) {
	// 需要加写锁，才能操作
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	c.m[key] = value
}

// Update 更新 key, 并且返回旧的 value 和是否更新成功
func (c *Cache) Update(key string, value interface{}) (interface{}, bool) {
	// 需要加写锁，才能操作

	c.rwLock.RLock()
	oldVal, ok := c.m[key]
	c.rwLock.RUnlock()
	if !ok {
		// 说明以前都没有这个 key
		return nil, false
	}

	c.rwLock.RLock()
	if val2 := c.m[key]; oldVal != val2 {
		return nil, false
	}
	c.rwLock.RUnlock()

	c.rwLock.Lock()
	c.m[key] = value
	c.rwLock.Unlock()

	return oldVal, true
}

// RandKey returns a random key from the cache
func (c *Cache) RandKey() string {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	keys := make([]string, 0, len(c.m))
	for k := range c.m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

// RandValue returns a random value
func RandValue() string {
	return fmt.Sprintf("value-%d", rand.Int())
}

func TestCache(t *testing.T) {
	// create a cache with 10 key-value pairs
	cache := NewCache(10)

	readers := 40
	writers := 10
	var wg sync.WaitGroup
	var hits, misses int
	rand.Seed(time.Now().UnixNano())

	// start readers
	for i := 0; i < readers; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			key := cache.RandKey()
			value, ok := cache.Get(key)
			if ok {
				fmt.Printf("reader-%d: hit, key = %s, value = %s\n", i, key, value)
				hits++
			} else {
				fmt.Printf("reader-%d: miss, key = %s\n", i, key)
				misses++
			}
			time.Sleep(time.Millisecond * 10)
		}()
	}

	// start writers
	for i := 0; i < writers; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			key := cache.RandKey()
			value := RandValue()
			cache.Put(key, value)
			fmt.Printf("writer-%d: update, key = %s, value = %s\n", i, key, value)
			time.Sleep(time.Millisecond * 10)
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the cache hit rate
	total := hits + misses
	rate := float64(hits) / float64(total) * 100
	fmt.Printf("cache hit rate: %d/%d (%.2f%%)\n", hits, total, rate)
}
