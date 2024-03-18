// @Author: Ciusyan 3/18/24

package cycle_1_3_18_3_22

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从根节点出发，查找 p和q 的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	// 从 Left 出发，查找 p和q 的最低公共祖先
	left := lowestCommonAncestor(root.Left, p, q)
	// 从 Right 出发，查找 p和q 的最低公共祖先
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		// 说明既存在左树，又存在右树，只能说明在 root
		return root
	}

	// 来到这里，至少有一边为 nil 了，返回对方，或者返回 nil
	if left == nil {
		return right
	}
	return left
}

/*
*
思路重复
对于给定的 p 和 q，去找它存在这颗树的公共祖先在什么地方，无非就这几种情况：
1. 在左子树（p 和 q 都在左子树）
2. 在右子树（p 和 q 都在右子树）
3. 在根节点（p 和 q 分布在左右）
4. 不存在 （p 或 q 不存在树上）

那么也就可以进行递归求解了，明确递归含义：
在 root 上，查找出 p 和 q 的最近公共祖先。
为什么呢？因为对于递归基的设定。是：
如果 root 为 nil 了，肯定得不到结果，如果 p 或 q 其中一个是 root，那么他们最近的公共祖先肯定是 root 了，

	因为它是此次调用中，最高的节点了，已经到 root 了

那么对于上面四种情况：
1.递归去左子树，肯定能得到结果，并且递归去右子树，一定不存在结果。
2.递归去右子树，肯定能得到结果，并且递归去左子树，一定不存在结果。
3.递归去左子树，能够得到结果，并且递归去右子树，也能得到结果。
4.去左右都获取不到结果
*/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}
