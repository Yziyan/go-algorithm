// @Author: Ciusyan 2023/7/27

package arrary

import (
	"fmt"
	"testing"
)

func TestMergeKSortArray(t *testing.T) {
	arrs := [][]int{
		{4, 6, 19},
		{5, 9, 13},
		{9, 10, 29},
	}

	array := mergeKSortedArray(arrs)
	fmt.Println(array)
}

func TestChangeSort(t *testing.T) {
	arr := []int{1, 4, 8, 2, 1, 5}

	changeSort(arr)

	fmt.Println(arr)

}
