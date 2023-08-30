// @Author: Ciusyan 2023/8/29

package binary_find

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExsit(t *testing.T) {
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 5, 6, 6}, 5))
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 5, 6, 6, 8}, 4))
	fmt.Println(Exsit([]int{1, 1, 2, 4, 4, 4, 6, 6}, 3))
	fmt.Println(Exsit([]int{1, 1, 2, 3, 4, 6, 6}, 3))
}

func TestNearestIndex(t *testing.T) {

	type args struct {
		arr []int
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{arr: []int{1, 2, 3}, num: 3}, want: 2},
		{name: "case2", args: args{arr: []int{1, 2, 2, 2, 3, 3, 4, 4, 5, 6}, num: 2}, want: 1},
		{name: "case3", args: args{arr: []int{1, 2, 2, 2, 3, 3, 4, 4, 6, 6, 6, 7}, num: 5}, want: 8},
		{name: "case4", args: args{arr: []int{1, 2}, num: 5}, want: -1},
		{name: "case5", args: args{arr: []int{1, 2}, num: 0}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NearestIndex(tt.args.arr, tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessIndex(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "case1", args: []int{1, 2, 3}, want: 0},
		{name: "case2", args: []int{10, 9, 8, 3, 3, 6, 5, 2}, want: 7},
		{name: "case3", args: []int{5, 2, 2, 2, 3, 1, 4, 4, 6, 6, 6, 7}, want: 5},
		{name: "case4", args: []int{3, 2, 4}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LessIndex(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
