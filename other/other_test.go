// @Author: Ciusyan 2023/7/23

package other

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	Print()
}

func Test1(t *testing.T) {
	s := "哈哈回去avax萨达"

	//for i, a := range s {
	//	fmt.Printf("%d %c\n", i, a)
	//}

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%d %c\n", i, runes[i])
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {

	LengthOfLongestSubstring("pwwkew")
}

func TestHttpClientPool(t *testing.T) {
	pool := NewHttpClientPool()

	for i := 0; i < 5; i++ {
		client := pool.Get()
		resp, err := client.Get("https://www.example.com\"")
		if err != nil {
			t.Error("失败了")

			return
		}

		_ = resp.Body.Close()
		pool.Put(client)
	}

}

func TestGoroutine(t *testing.T) {

	loopCount := 10000
	size := 10000

	type resp struct {
		age int
	}

	pushing := func(age int) *resp {
		// 模拟业务耗时
		return &resp{age: age}
	}

	goPushing := func(i int, ch chan<- *resp) {
		defer close(ch)
		for j := 0; j < size; j++ {
			res := pushing(i + j)
			ch <- res
		}
	}

	pushAfter := func(res *resp) {
		// 模拟业务耗时
	}

	testCases := []struct {
		name string

		execute func()
	}{
		// 会开过多的 Goroutine，协程调度开销很大、还需要涉及多个协程同步等待执行的结果
		{
			name: "【case1: for 循环，不协程】",
			execute: func() {
				wg := sync.WaitGroup{}
				for i := 0; i < loopCount; i++ {
					for j := 0; j < size; j++ {
						res := pushing(i + j)
						wg.Add(1)
						go func() {
							defer wg.Done()
							pushAfter(res)
						}()
					}
					wg.Wait()
				}
				fmt.Println()
			},
		},
		{
			// 这里每次只会去开一个协程，然后利用 channel，就避免了大量协程间的调度和等待问题。
			// 还有就是使用有 buffer 的 channel，可以减少协程频繁的阻塞和唤醒。
			name: "【case2:go 协程，channel 有 size 缓冲】",
			execute: func() {
				wg := sync.WaitGroup{}
				for i := 0; i < loopCount; i++ {
					ch := make(chan *resp, size)
					end := make(chan struct{})
					go func() {
						defer close(end)
						goPushing(i, ch)
						end <- struct{}{}
					}()

				OUTLOOP:
					for {
						select {
						case <-end:
							break OUTLOOP
						case res, ok := <-ch:
							if !ok {
								break
							}
							wg.Add(1)
							go func() {
								defer wg.Done()
								pushAfter(res)
							}()
						}
					}

					wg.Wait()
				}
			},
		},
		{
			// 这里每次只会去开一个协程，然后利用 channel，就避免了大量协程间的调度和等待问题。
			// 但是这里没有缓冲，所以按理来说要比上面慢
			name: "【case3:go 协程，channel 无缓冲】",
			execute: func() {
				wg := sync.WaitGroup{}
				for i := 0; i < loopCount; i++ {
					ch := make(chan *resp)
					end := make(chan struct{})
					go func() {
						defer close(end)
						goPushing(i, ch)
						end <- struct{}{}
					}()

				OUTLOOP:
					for {
						select {
						case <-end:
							break OUTLOOP
						case res, ok := <-ch:
							if !ok {
								break
							}
							wg.Add(1)
							go func() {
								defer wg.Done()
								pushAfter(res)
							}()
						}
					}

					wg.Wait()
				}
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			tt.execute()
			elapsed := time.Since(start)
			t.Logf("%s执行耗时: %s\n", tt.name, elapsed)
		})
	}

}

func TestNewConfig(t *testing.T) {
	// 使用的时候，就可以这样来使用
	// 这里使用必传参数
	cfg1 := NewConfig("localhost", 8080)

	// 这里可以使用可选项
	cfg2 := NewConfig("127.0.0.1", 433, WithTimeoutOption(5), WithIsLogOption(true))

	t.Logf("cfg1: %+v\n", cfg1)
	t.Logf("cfg2: %+v\n", cfg2)
}

func TestNewBuilder(t *testing.T) {
	builder := NewBuilder()

	p := builder.SetName("ciusyan").SetPrice(2.31).SetQuantity(3).Build()
	t.Logf("price: %f", p.GetPrice())

	// 假设中间获取 p 可能出错
	p1, err := builder.SetName("zhiyan").SetPrice(3.20).SetQuantity(3).BuildV1()
	if err != nil {
		t.Error(err)
		return
	}
	priceV1, err := p1.GetPriceV1()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("price: %f", priceV1)

	// 假设中间获取 p 可能出错，但是我们这里并不影响链式调用，
	// 可以将错误，交给最后一个链条处理，这样的代码会优雅很多。
	price, err := builder.SetName("zhiyan").SetPrice(0.00).SetQuantity(3).BuildV2().GetPrice()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%f", price)

	// 如果确保一定不会出错，出错就 panic
	p4 := builder.SetName("志颜").SetPrice(5.20).SetQuantity(100).BuildV2().MustProduct()
	t.Logf("%+v", p4)
}
