// @Author: Ciusyan 2023/9/2

package day_3

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	return getInfo(root).isBst
}

type info struct {
	isBst bool
	max   int
	min   int
}

func getInfo(root *TreeNode) *info {
	if root == nil {
		return nil
	}

	// 先搜集左右子树的信息
	lIo := getInfo(root.Left)
	rIo := getInfo(root.Right)

	// 填充 info 所需要的信息，先给他们个默认值，
	var (
		isBst = true
		max   = root.Val
		min   = root.Val
	)

	minF := func(v1, v2 int) int {
		if v1 > v2 {
			return v2
		}

		return v1
	}

	maxF := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}

		return v2
	}

	// 左边推
	if lIo != nil {
		max = maxF(max, lIo.max)
		min = minF(min, lIo.min)
		if !lIo.isBst || lIo.max >= root.Val {
			isBst = false
		}
	}

	// 右边也推
	if rIo != nil {
		max = maxF(max, rIo.max)
		min = minF(min, rIo.min)
		if !rIo.isBst || rIo.min <= root.Val {
			isBst = false
		}
	}

	// 返回 info
	return &info{
		isBst: isBst,
		max:   max,
		min:   min,
	}
}
