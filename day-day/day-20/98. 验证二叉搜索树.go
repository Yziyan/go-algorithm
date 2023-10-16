// @Author: Ciusyan 10/16/23

package day_20

// https://leetcode.cn/problems/validate-binary-search-tree/description/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 待收集的信息
	type info struct {
		isBst bool // 是否是二叉搜索树
		mx    int  // 最大值
		mn    int  // 最小值
	}

	// 递归函数
	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			// 这里不方便定义 info，所以交给上游处理
			return nil
		}

		// 先跟左右子树收集信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		// 利用左右子树构建 info 所需要的三个信息
		var (
			// 先默认是二叉搜索树
			isBst = true
			// 最大最小值先默认都是 root.Val
			mx = root.Val
			mn = root.Val
		)

		if leftInfo != nil {
			// 推大最大值
			if leftInfo.mx > mx {
				mx = leftInfo.mx
			}
			// 推小最小值
			if leftInfo.mn < mn {
				mn = leftInfo.mn
			}

			// 尝试查看这个是否是二叉搜索树
			if !leftInfo.isBst {
				// 左树都不是，肯定不是
				isBst = false
			}

			// 如果左树的最大值比 root.Val 还大，那么就不是
			if leftInfo.mx >= root.Val {
				isBst = false
			}
		}

		if rightInfo != nil {
			// 推大最大值
			if rightInfo.mx > mx {
				mx = rightInfo.mx
			}
			// 推小最小值
			if rightInfo.mn < mn {
				mn = rightInfo.mn
			}

			// 尝试查看这个是否是二叉搜索树
			if !rightInfo.isBst {
				// 右树都不是，肯定不是
				isBst = false
			}
			// 如果右树的最小值比 root.Val 还小，那么就不是
			if rightInfo.mn <= root.Val {
				isBst = false
			}
		}

		return &info{
			isBst: isBst,
			mx:    mx,
			mn:    mn,
		}
	}

	return getInfo(root).isBst
}
