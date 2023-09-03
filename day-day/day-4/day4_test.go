// @Author: Ciusyan 2023/9/3

package day_4

import (
	"fmt"
	"testing"
)

func TestMinDiff(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 48}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 49}

	fmt.Println(minDiffInBST(root))
}
