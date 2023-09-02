// @Author: Ciusyan 2023/9/2

package day_3

// https://leetcode.cn/problems/path-sum-ii/

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)

	// 从根节点开始搜索
	dfs(root, targetSum, &res, &[]int{})

	return res
}

func dfs(root *TreeNode, targetSum int, result *[][]int, path *[]int) {
	if root == nil {
		// 边界处理
		return
	}

	// 什么时候搜集结果呢？到达最后一层
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			// 来到这里，说明是一个结果，可以搜集，但是别忘记保存最后一个路径
			*result = append(*result, append(append([]int{}, *path...), root.Val))
		}

		return
	}

	// 保存路径
	*path = append(*path, root.Val)
	targetSum -= root.Val
	// 尝试去做右子树搜索
	dfs(root.Left, targetSum, result, path)
	dfs(root.Right, targetSum, result, path)

	// 还原现场，如果 path 传递的是切片的指针，那么就需要还原现场
	// 	如果 path 传递的是切片的引用，那么就不需要还原现场，因为有切片的扩容机制来兜底
	//	而且还原现场是返回当前层的时候还原就行了
	*path = (*path)[:len(*path)-1]
}
