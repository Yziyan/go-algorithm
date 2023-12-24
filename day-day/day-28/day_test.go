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

func TestGetNearLess(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want [][2]int
	}{
		{
			name: "Mixed values with duplicates",
			arr:  []int{4, 4, 1, 5, 6, 2, 7, 2},
			want: [][2]int{{-1, 2}, {-1, 2}, {-1, -1}, {2, 5}, {3, 5}, {2, -1}, {5, 7}, {2, -1}},
		},
		{
			name: "All elements same",
			arr:  []int{1, 1, 1, 1},
			want: [][2]int{{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}},
		},
		{
			name: "Ascending order",
			arr:  []int{1, 2, 3, 4, 5},
			want: [][2]int{{-1, -1}, {0, -1}, {1, -1}, {2, -1}, {3, -1}},
		},
		{
			name: "Descending order",
			arr:  []int{5, 4, 3, 2, 1},
			want: [][2]int{{-1, 1}, {-1, 2}, {-1, 3}, {-1, 4}, {-1, -1}},
		},
		// 更多测试用例可以添加在这里
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNearLess(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllTimesMinToMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Mixed values",
			arr:  []int{3, 1, 6, 4, 5, 2},
			want: 60,
		},
		{
			name: "Ascending order",
			arr:  []int{1, 2, 3, 4, 5},
			want: 36,
		},
		{
			name: "Descending order",
			arr:  []int{5, 4, 3, 2, 1},
			want: 36,
		},
		{
			name: "Contains duplicates",
			arr:  []int{1, 3, 2, 2, 1},
			want: 14,
		},
		// 更多测试用例可以添加在这里
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allTimesMinToMax(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaximalRectangle(t *testing.T) {
	testCases := []struct {
		name string
		args [][]byte

		want int
	}{
		{
			name: "Case1",
			args: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{
			name: "Case2",
			args: [][]byte{
				{'0'},
			},
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maximalRectangle(tc.args)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMatrixMul(t *testing.T) {

	mul := matrixMul([][]int{{1}, {1}}, [][]int{{1, 1}, {1, 0}})
	mul2 := matrixMul([][]int{{1, 1}, {1, 0}}, [][]int{{1}, {1}})
	t.Log(mul)
	t.Log(mul2)
}

func TestCowProblem(t *testing.T) {
	for i := 0; i < 20; i++ {
		got := cowProblem(i)
		got1 := cowProblem1(i)
		assert.Equal(t, got, got1)
	}

	n := 40
	got := cowProblem(n)
	got1 := cowProblem1(n)
	t.Log(got)
	t.Log(got1)
}
