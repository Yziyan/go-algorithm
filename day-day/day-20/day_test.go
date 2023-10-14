// @Author: Ciusyan 10/13/23

package day_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func TestIsCBT(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)

	assert.True(t, IsCBT(root))
}

func isAscendingOrder(nums []int) (string, bool) {
	if len(nums) == 0 {
		return "", true
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return "不是升序", false
		}
	}

	return "", true
}

func TestStackSort(t *testing.T) {

	testCases := []struct {
		name   string
		stackA Stack
	}{
		{
			name:   "【随机乱序】",
			stackA: []int{4, 1, 2, 4, 6, 2, 3, 5, 10, 8, 7, 9},
		},
		{
			name:   "【无元素】",
			stackA: []int{},
		},
		{
			name:   "【原本有序】",
			stackA: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			StackSort(tt.stackA)
			msg, order := isAscendingOrder(tt.stackA)
			assert.True(t, order, msg)
		})
	}
}
