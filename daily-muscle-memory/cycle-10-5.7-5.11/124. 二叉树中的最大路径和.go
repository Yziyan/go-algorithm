// @Author: Ciusyan 5/10/24

package cycle_10_5_7_5_11

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/

func maxPathSum(root *TreeNode) int {
	// 有无根：
	// 无根：1.max 在左子树，2.max 在右子树
	// 有根：1.左子树 + 根 + 右子树
	type info struct {
		maxPath         int
		maxPathFromRoot int // 从根出发的最大路径和
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			// 上游需要 check 对应的状态是否为有效的
			return nil
		}

		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			maxPath         = root.Val
			maxPathFromRoot = root.Val
		)

		if leftInfo != nil {
			maxPathFromRoot = max(maxPathFromRoot, root.Val+leftInfo.maxPathFromRoot)
			maxPath = max(maxPath, leftInfo.maxPath)
		}

		if rightInfo != nil {
			maxPathFromRoot = max(maxPathFromRoot, root.Val+rightInfo.maxPathFromRoot)
			maxPath = max(maxPath, rightInfo.maxPath)
		}

		maxPath = max(maxPath, maxPathFromRoot)

		// 左 + 根 + 右
		if leftInfo != nil && rightInfo != nil {
			if leftInfo.maxPathFromRoot > 0 && rightInfo.maxPathFromRoot > 0 {
				maxPath = max(maxPath, leftInfo.maxPathFromRoot+rightInfo.maxPathFromRoot+root.Val)
			}
		}

		// 那么当前节点的最大路径和，要么是左子树，要么是右子树，要么是包含根节点的结果

		return &info{
			maxPath:         maxPath,
			maxPathFromRoot: maxPathFromRoot,
		}
	}

	if root == nil {
		return 0
	}

	return getInfo(root).maxPath
}
