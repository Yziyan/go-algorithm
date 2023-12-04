// @Author: Ciusyan 12/4/23

package day_26

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLivePossibility(t *testing.T) {
	tests := []struct {
		name string
		row  int
		col  int
		k    int
		n    int
		m    int
		want float64
	}{
		{
			name: "Test1",
			row:  1,
			col:  1,
			k:    2,
			n:    3,
			m:    3,
			want: 0.75,
		},
		{
			name: "Test2",
			row:  0,
			col:  0,
			k:    3,
			n:    2,
			m:    2,
			want: 0.125,
		},
		{
			name: "Test3",
			row:  2,
			col:  2,
			k:    4,
			n:    5,
			m:    5,
			want: 0.84375,
		},
		// 可以添加更多测试用例...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := livePossibility(tc.row, tc.col, tc.k, tc.n, tc.m)
			assert.Equal(t, tc.want, got)
		})
	}
}
