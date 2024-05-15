// @Author: Ciusyan 5/10/24

package cycle_10_5_7_5_11

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/

/**
思路重复：
要找到最大路径和，可以使用二叉树的递归套路来解：
首先思考，根节点有无影响结果吗？其实是影响的，因为最大路径可能有下面几种情况：
1.就是根节点
2.在左子树
3.在右子树
4.根+左子树
5.根+右子树
6.根+左+右

那么可以总结出需要两个元信息：最大路径和、从根节点出发的最大路径和

然后呢？先抓左右子树的信息，最大路径和就是上面的情况。
他俩信息默认都设置为 root.Val，如果左子树、右能抓到信息，尝试更新两个信息，
如果左右子树都能抓到信息，再尝试更新总的信息。
值得一提的是，当为 nil 的情况，不能直接返回两个值为 0 的 info 信息，因为后续需要判断，这个值是否是有效的
*/

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
