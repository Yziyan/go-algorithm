// @Author: Ciusyan 5/23/24

package cycle_12_5_22_5_26

// https://leetcode.cn/problems/validate-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路重复：
二叉搜索树有一个基本的性质就是：中序遍历是有序的。
所以我们只需要验证，中序遍历，是否是有序的即可，每次记录下前一个遍历的节点即可。

但在这里，我会使用二叉树的递归套路来编写：
收集信息：
1.是否是二叉搜索树
2.树的最大值
3.树的最小值

当有了这三个信息后，我们就可以去判断：当前 root 是否是 BST 了，
只需要满足：
1.左子树是二叉搜索树
2.右子树是二叉搜索树
3.左子树的最大值小于 root.Val
3.右子树的最小值大于 root.Val

就说明这棵树是二叉搜索树。
*/

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
