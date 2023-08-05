// @Author: Ciusyan 2023/8/4

package tree

// https://leetcode.cn/problems/path-sum-ii/

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	// 轨迹
	track := make([]int, 0)
	// 从根节点开始搜索
	pathSumDfs(root, targetSum, track, &res)

	return res
}

// 从 root 开始搜索，
func pathSumDfs(root *TreeNode, sum int, track []int, results *[][]int) {
	if root.Left == nil && root.Right == nil {
		// 如果搜索到了叶子节点，那么看看是否需要记录结果
		if sum == root.Val {
			// 别忘记添加最后一个节点的值
			*results = append(*results, append(append([]int{}, track...), root.Val))
		}

		return
	}

	// 来到这里，记录轨迹，并且减去对应的值
	sum -= root.Val
	track = append(track, root.Val)
	if root.Left != nil {
		// 往左子树开始搜索
		pathSumDfs(root.Left, sum, track, results)
	}

	if root.Right != nil {
		// 往右子树搜索
		pathSumDfs(root.Right, sum, track, results)
	}

	// 最好还是还原一下现场，但其实我们这里不还原现场，
	// 也有 Go 切片的扩容机制给我们兜底，因为扩容前后，
	// 底层引用的数组就不在相同了，那么对后续的操作，也不会影响到旧切片的值
	track = track[:len(track)-1]
}
