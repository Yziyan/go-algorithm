// @Author: Ciusyan 11/1/23

package other

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
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
		// 并发环境，share <= 200
		assert.Equal(t, 200, share)
		t.Logf("执行后的值: %d", share)
	})
}

func TestRwMutex(t *testing.T) {
	var rwMu sync.RWMutex

	rwMu.Lock()

}
