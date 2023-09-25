// @Author: Ciusyan 9/26/23

package day_17

import "testing"

func TestCountingSort(t *testing.T) {
	nums := []int{4, 2, 7, 35, 29, -1, -6, 4, 3, 18, 34, -10}
	CountingSort(nums)
	t.Log(nums)
}
