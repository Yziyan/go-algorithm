// @Author: Ciusyan 12/12/23

package day_27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleQueue(t *testing.T) {
	queue := NewDoubleQueue()
	assert.True(t, queue.Size() == 0)
	queue.OfferRight(3)
	queue.OfferRight(4)
	queue.OfferRight(5)
	assert.True(t, queue.Size() == 3)
	assert.True(t, queue.PollLeft() == 3)
	assert.True(t, queue.PollRight() == 5)
	assert.True(t, queue.Size() == 1)
	queue.OfferLeft(10)
	queue.OfferLeft(20)
	assert.True(t, queue.PollLeft() == 20)
	assert.True(t, queue.Size() == 2)
}

func TestSlidingWindowMaxArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Normal case",
			arr:  []int{1, 3, -1, -3, 5, 3, 6, 7},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "Decreasing elements",
			arr:  []int{7, 6, 5, 4, 3, 2, 1},
			want: []int{7, 6, 5, 4, 3},
		},
		{
			name: "Increasing elements",
			arr:  []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{3, 4, 5, 6, 7},
		},
		{
			name: "All same elements",
			arr:  []int{2, 2, 2, 2, 2, 2},
			want: []int{2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slidingWindowMaxArray(tt.arr, 3)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllLessNumSubArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		num  int
		want int
	}{
		{"Basic case", []int{1, 3, 5, 7, 9}, 2, 9},
		{"Single element", []int{5}, 0, 1},
		{"All elements same", []int{2, 2, 2}, 0, 6},
		{"No valid subarray", []int{10, 20, 30}, 5, 3},
		{"Mixed values", []int{1, 2, 3, 4, 5}, 3, 14},
		{"Empty array", []int{}, 3, 0},
		{"Large range", []int{1, 2, 3, 100, 101}, 100, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allLessNumSubArray(tt.arr, tt.num)
			assert.Equal(t, tt.want, got)
		})
	}
}
