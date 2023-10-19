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

func TestMaxSubBSTTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    *TreeNode
		expected int
	}{
		{
			name:     "【单一节点，树中只有一个节点】",
			input:    &TreeNode{Val: 5},
			expected: 1,
		},
		{
			name: "【完全平衡的二叉搜索树（BST）】",
			input: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 15},
			},
			expected: 3,
		},
		{
			name: "【一个完整的非BST树】",
			input: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 5},
			},
			expected: 1,
		},
		{
			name: "【非平衡树，但有BST子树】",
			input: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{
					Val:  15,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: 2,
		},
		{
			name: "【更复杂的树，混合了BST和非BST部分】",
			input: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 15},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 10},
					},
				},
			},
			expected: 5,
		},
		// 添加更多测试用例
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSubBSTTree(tt.input)
			assert.Equal(t, tt.expected, got)
		})

	}
}
