// @Author: Ciusyan 4/2/24

package other

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

// 结构体用于存储HTTP请求的结果
type httpResult struct {
	url      string
	status   string
	duration time.Duration
}

func TestHttpChannel(t *testing.T) {
	urls := []string{
		"https://www.nowcoder.com/",
		"http://juejin.com",
		"http://baidu.com",
		// 添加更多URLs
	}

	// 创建一个通道，用于传递httpResult
	resultsCh := make(chan httpResult, len(urls))

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(url)
			duration := time.Since(start)
			if err != nil {
				fmt.Println("Error:", err)
				resultsCh <- httpResult{url: url, status: "Failed", duration: duration}
				return
			}
			resultsCh <- httpResult{url: url, status: resp.Status, duration: duration}
			resp.Body.Close()
		}(url)
	}

	// 关闭通道，在所有协程完成后
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// 从通道中读取并打印结果
	for result := range resultsCh {
		fmt.Printf("URL: %s, Status: %s, Duration: %v\n", result.url, result.status, result.duration)
	}
}
