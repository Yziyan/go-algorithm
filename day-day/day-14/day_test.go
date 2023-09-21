// @Author: Ciusyan 9/21/23

package day_14

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{5, 1, 1, 2, 0, 0}
	// 4, 2, 3 5 5 5 9 7 6 8
	fmt.Println(sortArray(nums))
}
