// @Author: Ciusyan 10/8/23

package day_19

import (
	"fmt"
	"testing"
)

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func TestCodec(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Right.Left = NewTreeNode(4)
	root.Right.Right = NewTreeNode(5)

	codec := Constructor()
	serialize := codec.serialize(root)
	t.Log(serialize)
	root = codec.deserialize(serialize)
	fmt.Println()
}
