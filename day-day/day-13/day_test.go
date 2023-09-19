// @Author: Ciusyan 9/19/23

package day_13

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	nums := []int{5, 4, 9, 5, 2, 5, 5, 7, 6, 8, 3}
	// 4, 2, 3 5 5 5 9 7 6 8
	fmt.Println(getPivotPoint(nums, 0, len(nums)))
}
