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

func TestOddTimesNum1(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{1, 2, 2}, want: 1},
		{name: "case2", args: []int{10, 9, 10, 3, 3, 9, 2, 1, 2, 1, 2}, want: 2},
		{name: "case3", args: []int{3, 2, 4, 2, 2, 3, 2, 4, 4, 6, 4, 4, 4}, want: 6},
		{name: "case4", args: []int{3, 2, 4, 2, 4}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OddTimesNum1(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOddTimesNum2(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{1, 2, 2, 2}, want: 3},
		{name: "case2", args: []int{10, 9, 4, 10, 3, 3, 9, 2, 1, 2, 1, 2}, want: 6},
		{name: "case3", args: []int{3, 2, 4, 2, 2, 4, 3, 2, 4, 4, 6, 4, 4, 4}, want: 10},
		{name: "case4", args: []int{3, 2, 4, 2, 3, 4, 1, 6}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OddTimesNum2(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
