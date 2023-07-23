// @Author: Ciusyan 2023/7/23

package other_test

import (
	"fmt"
	"github.com/Yziyan/go-algorithm/other"
	"testing"
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
