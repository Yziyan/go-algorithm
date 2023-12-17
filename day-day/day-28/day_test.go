// @Author: Ciusyan 12/17/23

package day_28

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNearLessNoRepeat(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want [][2]int
	}{
		{
			name: "Test Case 1",
			arr:  []int{3, 4, 1, 5, 6, 2, 7},
			want: [][2]int{{-1, 2}, {0, 2}, {-1, -1}, {2, 5}, {3, 5}, {2, -1}, {5, -1}},
		},
		{
			name: "Test Case 2",
			arr:  []int{1, 2, 3, 4, 5},
			want: [][2]int{{-1, -1}, {0, -1}, {1, -1}, {2, -1}, {3, -1}},
		},
		// 更多测试用例...
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getNearLessNoRepeat(tc.arr)
			assert.Equal(t, tc.want, got)
		})
	}
}
