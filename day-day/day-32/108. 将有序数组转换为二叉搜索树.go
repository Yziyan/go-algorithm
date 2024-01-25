// @Author: Ciusyan 1/25/24

package day_32

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	// 准备一个递归函数，含义是：
	// 将 [l, r) 位置，转换出一颗平衡二叉搜索树，返回 root 节点
	var process func(nums []int, l, r int) *TreeNode
	process = func(nums []int, l, r int) *TreeNode {
		if l == r {
			return nil
		}

		// 取终点
		mid := l + ((r - l) >> 1)
		// 先构建出根节点
		root := &TreeNode{Val: nums[mid]}
		// 再去 [l, mid) 上构建出左子树
		root.Left = process(nums, l, mid)
		// 和去 [mid+1, r) 上构建出右子树
		root.Right = process(nums, mid+1, r)

		return root
	}

	// 去 [0, n) 上构建二叉树
	return process(nums, 0, len(nums))
}
