// @Author: Ciusyan 9/14/23

package day_9

import (
	"fmt"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	arr := []int{2, 4, 1, 6, 2, 1, 3, 8, 5, 6, 3, 10}
	MergeSort2(arr)
	fmt.Println(arr)
}
