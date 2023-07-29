// @Author: Ciusyan 2023/7/23

package other

import (
	"fmt"
	"sync"
)

// 最终实现的效果形如：AB12CD34EF......
func printNum(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 26; i++ {
		fmt.Print(i)
		// 通知另一个协程打印字符串
		ch <- struct{}{}

		// 等待另一个协程通知了再继续打印下一个数字
		<-ch
	}
}

func printStr(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 'A'; i <= 'Z'; i++ {
		// 等待通知了再打印
		<-ch
		fmt.Printf("%c", i)

		// 通知另一个协程打印数字
		ch <- struct{}{}
	}
}

func Print() {
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go printNum(ch, &wg)
	go printStr(ch, &wg)

	wg.Wait()
}
