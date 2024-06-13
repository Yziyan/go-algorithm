// @Author: Ciusyan 6/4/24

package cycle_14_5_27_5_31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路重复：
对于这一棵树，如何是相同的呢？其实判定规则挺简单的：
1.根节点的值相同
2.左子树是相同的树
3.右子树是相同的数

那么就只需要递归判断即可，那么核心点就是找到递归基。
1. 当两颗树都是 nil 的时候，相同
2. 当两棵树只有一颗是 nil 的时候，不同
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
