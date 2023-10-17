// @Author: Ciusyan 2023/7/23

package other_test

import (
	"context"
	"fmt"
	"github.com/Yziyan/go-algorithm/other"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	other.Print()
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

	other.LengthOfLongestSubstring("pwwkew")
}

func TestHttpClientPool(t *testing.T) {
	pool := other.NewHttpClientPool()

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

func Busyness(ctx context.Context) {
	time.Sleep(3 * time.Second)
	fmt.Println("业务完成了")
}
