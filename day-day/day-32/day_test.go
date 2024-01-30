// @Author: Ciusyan 1/23/24

package day_32

import (
	"math"
	"testing"
)

func TestName(t *testing.T) {
	s := "*"

	t.Log(len(s))
	t.Log(string(s[0]))
	t.Log(math.MaxInt)

	t.Log(!true)
	t.Log(!false)

	t.Logf("256TB = %dMB", 256*1024*1024)
	t.Logf("2^22 * 64MB = %dMB", int(math.Pow(2, 22)*64))
}

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	sum := maxPathSum(root)
	t.Log(sum)
}
