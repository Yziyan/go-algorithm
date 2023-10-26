// @Author: Ciusyan 10/26/23

package other

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
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
