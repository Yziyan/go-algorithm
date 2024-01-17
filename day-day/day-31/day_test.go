// @Author: Ciusyan 1/17/24

package day_31

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{1, 0, 1, 1, 1}

	b := search(nums, 0)
	t.Log(b)
}
