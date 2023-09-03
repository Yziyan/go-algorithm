// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}

	return buildTree(nums, 0, len(nums))
}

// 从 [begin, end) 构建二叉搜索树
func buildTree(nums []int, begin, end int) *TreeNode {
	if end-begin < 1 {
		return nil
	}

	// 尽量让树平衡
	mid := begin + (end-begin)>>1
	root := &TreeNode{Val: nums[mid]}
	// 递归构建左右子树
	root.Left = buildTree(nums, begin, mid)
	root.Right = buildTree(nums, mid+1, end)

	return root
}
