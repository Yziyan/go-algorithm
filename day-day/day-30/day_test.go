// @Author: Ciusyan 1/2/24

package day_30

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubArrModM(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		m    int
		want int
	}{
		{
			name: "Test 1",
			arr:  []int{3, 1, 4, 2},
			m:    6,
			want: 5,
		},
		{
			name: "Test 2",
			arr:  []int{1, 2, 3},
			m:    4,
			want: 3,
		},
		{
			name: "Test 3",
			arr:  []int{1, 2, 3, 4},
			m:    5,
			want: 4,
		},
		{
			name: "Test 4",
			arr:  []int{10, 12, 7},
			m:    8,
			want: 7,
		},
		{
			name: "Test 5",
			arr:  []int{6, 7, 9},
			m:    5,
			want: 4,
		},
		{
			name: "Large Array Test 1",
			arr:  []int{12, 3, 14, 56, 77, 13, 4, 25, 33, 45, 67, 89, 2, 6, 8},
			m:    10,
			want: 9,
		},
		{
			name: "Large Array Test 2",
			arr:  []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
			m:    15,
			want: 14,
		},
		{
			name: "Large Array with Small m",
			arr:  []int{7, 11, 5, 2, 8, 13, 21, 1, 3},
			m:    3,
			want: 2,
		},
		{
			name: "Large Values",
			arr:  []int{102, 304, 506, 708, 910, 111, 313, 515},
			m:    50,
			want: 49,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := subArrModM(tc.arr, tc.m)
			assert.Equal(t, tc.want, got)
		})
	}

}
