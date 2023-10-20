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
	// 构建测试用例表格
	testCases := []struct {
		name     string
		input    *TreeNode
		expected bool
	}{
		{
			name:     "Empty Tree",
			input:    nil,
			expected: true,
		},
		{
			name: "Complete Binary Tree",
			input: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "Incomplete Binary Tree",
			input: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			name: "Full Binary Tree",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: true,
		},
		{
			name: "Not Full Binary Tree",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 6},
				},
			},
			expected: false,
		},
		{
			name: "测试用例 3：满二叉树，应该返回 true",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "测试用例 4：完全二叉树，但不是满二叉树，应该返回 true",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: true,
		},
		{
			name: "测试用例 5：非完全二叉树，应该返回 false",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: false,
		},
	}

	// 遍历测试用例并运行测试
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isCompleteTree(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
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
		args     *TreeNode
		expected int
	}{
		{
			name:     "【单一节点，树中只有一个节点】",
			args:     &TreeNode{Val: 5},
			expected: 1,
		},
		{
			name: "【完全平衡的二叉搜索树（BST）】",
			args: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 15},
			},
			expected: 3,
		},
		{
			name: "【一个完整的非BST树】",
			args: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 5},
			},
			expected: 1,
		},
		{
			name: "【非平衡树，但有BST子树】",
			args: &TreeNode{
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
			args: &TreeNode{
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
			got := MaxSubBSTTree(tt.args)
			assert.Equal(t, tt.expected, got)
		})
	}
}
