// @Author: Ciusyan 2023/8/30

package bit_operation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	type args struct {
		arr  []int
		i, j int
	}

	verfiy := func(arg args) bool {
		vi := arg.arr[arg.i]
		vj := arg.arr[arg.j]
		Swap(arg.arr, arg.i, arg.j)

		if vi == arg.arr[arg.j] && vj == arg.arr[arg.i] {
			return true
		}

		return false
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{arr: []int{1, 2, 3}, i: 0, j: 0}, want: true},
		{name: "case1", args: args{arr: []int{1, 2, 3}, i: 0, j: 2}, want: true},
		{name: "case1", args: args{arr: []int{1, 9, 3, 5}, i: 1, j: 2}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := verfiy(tt.args)
			assert.Equal(t, got, tt.want)
		})
	}
}
