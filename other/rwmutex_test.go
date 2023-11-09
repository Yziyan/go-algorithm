// @Author: Ciusyan 11/1/23

package other

import (
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
