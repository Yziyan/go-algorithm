// @Author: Ciusyan 10/26/23

package other

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	// 设置最大并发数
	maxConcurrent := 4

	// 创建一个信号量
	sem := semaphore.NewWeighted(int64(maxConcurrent))

	// 创建等待组，以确保所有协程完成
	var wg sync.WaitGroup

	// 创建一个上下文
	ctx := context.Background()

	// 任务数量
	numTasks := 10

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()

			// 请求信号量，控制并发数
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Printf("任务 %d 不能获取信号量锁: %v\n", taskID, err)
				return
			}

			// 执行任务
			fmt.Printf("任务 %d 正在执行\n", taskID)
			time.Sleep(2 * time.Second)

			// 释放信号量
			sem.Release(1)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()
}

var maxConcurrent int32 = 1

func TestSemaphore2(t *testing.T) {

	// 使用信号量锁
	sem := semaphore.NewWeighted(int64(atomic.LoadInt32(&maxConcurrent)))

	wg := sync.WaitGroup{}
	// ctx := context.Background()

	go func() {
		time.Sleep(1 * time.Second)
		updateMaxConcurrent(5)
	}()

	for i := 1; i <= 10; i++ {

		wg.Add(1)

		go func(taskId int) {
			// 别忘了释放资源
			defer wg.Done()

			if !tryAcquireWithMaxConcurrent(sem) {
				// 没获取到锁
				t.Logf("TaskId = %d 尚未获取到锁", taskId)
				return
			}
			// 这里说明成功获取锁了，退出时记得释放
			defer sem.Release(1)

			// 执行任务
			t.Logf("TaskId = %d 正在执行中", taskId)
			time.Sleep(2 * time.Second)

		}(i)

		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}

func tryAcquireWithMaxConcurrent(sem *semaphore.Weighted) bool {
	max := atomic.LoadInt32(&maxConcurrent)
	if max <= 0 {
		return false
	}
	if sem.TryAcquire(1) {
		atomic.AddInt32(&maxConcurrent, -1)
		return true
	}
	return false
}

// 在运行时更新 maxConcurrent 的值
func updateMaxConcurrent(newMaxConcurrent int32) {
	atomic.StoreInt32(&maxConcurrent, newMaxConcurrent)
}
