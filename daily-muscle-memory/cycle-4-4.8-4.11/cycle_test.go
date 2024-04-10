// @Author: Ciusyan 4/10/24

package cycle_4_4_8_4_11

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	m := make(map[int]int, 1)

	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(500 * time.Millisecond)
			m[i] = i
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
			time.Sleep(500 * time.Millisecond)
		}
	}()

	select {}
}
