// @Author: Ciusyan 5/7/24

package cycle_10_5_7_5_11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(nums []int, l, r int) *TreeNode
	build = func(nums []int, l, r int) *TreeNode {
		if r-l < 1 {
			return nil
		}

		mid := l + ((r - l) >> 1)
		root := &TreeNode{Val: nums[mid]}
		root.Left = build(nums, l, mid)
		root.Right = build(nums, mid+1, r)
		return root
	}

	return build(nums, 0, len(nums))
}
