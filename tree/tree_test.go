// @Author: Ciusyan 2023/7/27

package tree

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T) {

	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Right = NewTreeNode(5)
	root.Right.Right = NewTreeNode(4)

	res := RightSideView(root)

	for _, v := range res {
		fmt.Println(v)
	}
}

func TestMinDiffInBST(t *testing.T) {

	root := NewTreeNode(1)
	root.Left = NewTreeNode(0)
	root.Right = NewTreeNode(48)
	root.Right.Left = NewTreeNode(12)
	root.Right.Right = NewTreeNode(49)

	bst := MinDiffInBST(root)
	fmt.Println(bst)
}