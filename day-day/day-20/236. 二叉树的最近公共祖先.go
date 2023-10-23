// @Author: Ciusyan 10/23/23

package day_20

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 定义 info
	type info struct {
		lca     *TreeNode
		isFindP bool
		isFindQ bool
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			// 空树好设置
			return &info{}
		}

		// 收集左右子树信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			lca     *TreeNode
			isFindP bool
			isFindQ bool
		)

		if leftInfo.isFindP || rightInfo.isFindP || root == p {
			// 说明左右子树有一个存在 P，那么当前 root 肯定存在，或者就是当前节点本身就等于 P
			isFindP = true
		}

		if leftInfo.isFindQ || rightInfo.isFindQ || root == q {
			// 说明左右子树有一个存在 Q，那么当前 root 肯定存在，或者就是当前节点本身就等于 Q
			isFindQ = true
		}

		// lca 如何构建呢？
		// lca 存在于左右子树
		if leftInfo.lca != nil {
			lca = leftInfo.lca
		} else if rightInfo.lca != nil {
			lca = rightInfo.lca
		} else if isFindP && isFindQ {
			// 这个说明左右子树一边发现一个节点
			lca = root
		}

		// 利用从左右子树上收集的信息构建出 root 所需要的信息
		return &info{
			lca:     lca,
			isFindP: isFindP,
			isFindQ: isFindQ,
		}
	}

	return getInfo(root).lca
}

// 方法一：直接递归来求解
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 如果某个节点和是 root，那么不用看了，p 和 q 的公共祖先肯定是 root
		return root
	}

	// 否则来看看左右子树的公共祖先是谁
	leftLca := lowestCommonAncestor1(root.Left, p, q)
	rightLca := lowestCommonAncestor1(root.Right, p, q)

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
