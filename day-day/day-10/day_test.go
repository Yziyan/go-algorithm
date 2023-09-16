// @Author: Ciusyan 9/16/23

package day_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallSum(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{4, 2, 1, 6, 3}, want: 10},
		{name: "case2", args: []int{2, 2, 1, 6, 3, 2, 4}, want: 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SmallSum(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
