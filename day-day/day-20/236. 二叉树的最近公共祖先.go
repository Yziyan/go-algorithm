// @Author: Ciusyan 10/23/23

package day_20

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

// 方法一：直接递归来求解
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 如果某个节点和是 root，那么不用看了，p 和 q 的公共祖先肯定是 root
		return root
	}

	// 否则来看看左右子树的公共祖先是谁
	leftLca := lowestCommonAncestor(root.Left, p, q)
	rightLca := lowestCommonAncestor(root.Right, p, q)

	// 如果两者都有值 nil，说明是 root
	if leftLca != nil && rightLca != nil {
		return root
	}

	// 来到这里有三种情况：
	// 1.左子树的LCA 2.右子树的LCA 3. 没有

	// 这就说明在左子树能找到最近公共祖先
	if leftLca != nil {
		return leftLca
	}

	// 在这里就说明找不到，得看看右子树有没有，直接返回：1.在右子树 2.不存在
	return rightLca
}
