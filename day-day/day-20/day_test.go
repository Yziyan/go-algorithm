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
