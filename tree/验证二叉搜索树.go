// @Author: Ciusyan 2023/8/7

package tree

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return getBstInfo(root).isBst
}

// Info 信息
type bstInfo struct {
	isBst bool
	max   int
	min   int
}

func newBstInfo(isBst bool, max, min int) *bstInfo {
	return &bstInfo{isBst: isBst, max: max, min: min}
}

func getBstInfo(root *TreeNode) *bstInfo {
	if root == nil {
		return nil
	}

	// 先获取左右子树的信息
	leftInfo := getBstInfo(root.Left)
	rightInfo := getBstInfo(root.Right)

	isBst := true
	mx := root.Val
	mn := root.Val

	// 左子树的信息不为 nil
	if leftInfo != nil {
		// 尝试更新最大最小值
		mx = max(mx, leftInfo.max)
		mn = min(mn, leftInfo.min)
	}

	// 右子树的信息不为 nil
	if rightInfo != nil {
		// 尝试更新最大最小值
		mx = max(mx, rightInfo.max)
		mn = min(mn, rightInfo.min)
	}

	// 看看能否推出不是二叉搜索树
	if (leftInfo != nil && !leftInfo.isBst) || (rightInfo != nil && !rightInfo.isBst) {
		// 说明左子树或者右子树本身就不是二叉搜索树了
		isBst = false
	}

	if (leftInfo != nil && leftInfo.max >= root.Val) || (rightInfo != nil && rightInfo.min <= root.Val) {
		// 说明左右子树的大小关系不对
		isBst = false
	}

	return newBstInfo(isBst, mx, mn)
}
