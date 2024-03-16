// @Author: Ciusyan 3/16/24

package phase_1

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

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

		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			mx    = root.Val
			mn    = root.Val
			isBst = true
		)

		if leftInfo != nil {
			// 左子树能搜集信息，尝试更新最大最小值
			mx = max(mx, leftInfo.mx)
			mn = min(mn, leftInfo.mn)
		}

		if rightInfo != nil {
			// 右子树能搜集信息，尝试更新最大最小值
			mx = max(mx, rightInfo.mx)
			mn = min(mn, rightInfo.mn)
		}

		if (leftInfo != nil && !leftInfo.isBst) || (rightInfo != nil && !rightInfo.isBst) {
			isBst = false
		}

		if (leftInfo != nil && root.Val <= leftInfo.mx) || (rightInfo != nil && root.Val >= rightInfo.mn) {
			isBst = false
		}

		return &info{
			mx:    mx,
			mn:    mn,
			isBst: isBst,
		}
	}

	return getInfo(root).isBst
}
