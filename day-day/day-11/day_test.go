// @Author: Ciusyan 9/16/23

package day_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePairs(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{4, 2, 1, 6, 3}, want: 1},
		{name: "case2", args: []int{2, 5, 1, 6, 1, 2, 4}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reversePairs(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
