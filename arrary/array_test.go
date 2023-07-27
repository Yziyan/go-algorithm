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
