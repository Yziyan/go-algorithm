// @Author: Ciusyan 3/8/24

package day_34

import "testing"

func TestFindMissingRanges(t *testing.T) {

	ranges := findMissingRanges([]int{0, 2, 3, 10, 50, 78}, -3, 99)

	t.Log(ranges)
}
