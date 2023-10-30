// @Author: Ciusyan 10/30/23

package other

import (
	"sync"
	"testing"
	"time"
)

func TestMiniSema(t *testing.T) {
	// 初始设置为 1
	sema := NewMiniSema()
	// 如果需要修改，可以这样传入
	// sema := NewMiniSema(WithMaxConcurrentOption(3))

	updateFunc := func(cm int64, target int32) {
		time.Sleep(time.Duration(cm) * time.Second)
		sema.UpdateMaxConcurrent(target)
		t.Logf("max修改为: %d, update: %#v", target, sema)
	}

	// 模拟修改成 3
	go updateFunc(2, 3)

	// 模拟修改成 1
	go updateFunc(6, 1)

	// 模拟改成 5
	go updateFunc(10, 5)

	// 模拟改成 2
	go updateFunc(12, 2)

	business(t, sema)
}

func business(t *testing.T, sema *MiniSema) {

	var wg sync.WaitGroup

	for i := 1; i <= 150; i++ {
		wg.Add(1)
		go func(taskId int) {
			defer wg.Done()

			if !sema.TryAcquire() {
				// 没获取到锁
				t.Logf("TaskId = %d 尚未获取到锁", taskId)
				return
			}
			// 获取信号量
			defer func() {
				sema.Release(1)
				t.Logf("Release: %#v", sema)
			}() // 释放信号量

			// 执行任务
			t.Logf("TaskID = %d 开始执行, Acquire: %#v\n", taskId, sema)
			time.Sleep(2 * time.Second)
			t.Logf("TaskID = %d 执行结束\n", taskId)
		}(i)

		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

func TestMiniSema_ListenMaxConcurrent(t *testing.T) {

	maxSize := []int32{1, 3, 6, 2, 5, 1}
	idx := 0
	count := 0
	// 新建一个一开始能有 3 个并发数量的 sema
	sema := NewMiniSema(WithMaxConcurrentOption(3))

	// 假设 2 秒改一次并发数
	sema.ListenMaxConcurrent(func() int32 {
		count++
		if (count & 1) != 1 {
			// 如果是偶数次，那当做没有改过，复用上一次的
			return maxSize[idx]
		}

		if idx >= len(maxSize) {
			idx = 0
		}

		val := maxSize[idx]
		idx++
		return val
	}, 2*time.Second)

	business(t, sema)
}
