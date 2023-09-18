// @Author: Ciusyan 9/18/23

package day_12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNationalFlag(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{4, 2, 1, 6, 3}, want: 4},
		{name: "case2", args: []int{2, 5, 1, 6, 4, 2, 4}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NationalFlag(tt.args, 0, len(tt.args)-1)
			assert.Equal(t, tt.want, got[0]+got[1])
		})
	}
}
