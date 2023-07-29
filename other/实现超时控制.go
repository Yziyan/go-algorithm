// @Author: Ciusyan 2023/7/29

package other

import (
	"context"
	"fmt"
	"time"
)

func Control() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	end := make(chan struct{})

	go func() {
		Busyness(ctx)
		// 业务做完了，给一个信号
		end <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("执行业务超时了")
	case <-end:
		fmt.Println("业务处理完成")
	}

}

func Busyness(ctx context.Context) {
	time.Sleep(2 * time.Second)

	ctx.Done()

	fmt.Println("执行业务方法")
}
