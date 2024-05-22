// @Author: Ciusyan 5/23/24

package cycle_12_5_22_5_26

// https://leetcode.cn/problems/validate-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	type info struct {
		mx    int
		mn    int
		isBst bool
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			return nil
		}

		l, r := getInfo(root.Left), getInfo(root.Right)
		var (
			mx    = root.Val
			mn    = root.Val
			isBst = true
		)

		if l != nil {
			mx = max(mx, l.mx)
			mn = min(mn, l.mn)
			if !l.isBst || l.mx >= root.Val {
				isBst = false
			}
		}

		if r != nil {
			mx = max(mx, r.mx)
			mn = min(mn, r.mn)
			if !r.isBst || r.mn <= root.Val {
				isBst = false
			}
		}

		return &info{
			mx:    mx,
			mn:    mn,
			isBst: isBst,
		}
	}

	if root == nil {
		return true
	}

	return getInfo(root).isBst
}
