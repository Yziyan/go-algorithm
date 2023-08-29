// @Author: Ciusyan 2023/8/29

package binary_find

import (
	"fmt"
	"testing"
)

func TestExsit(t *testing.T) {
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 5, 6, 6}, 5))
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 5, 6, 6, 8}, 4))
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 6, 6}, 3))
	fmt.Println(Exsit([]int{1, 1, 2, 3, 4, 6, 6}, 3))
}
