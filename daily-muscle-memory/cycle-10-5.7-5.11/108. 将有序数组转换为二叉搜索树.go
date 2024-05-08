// @Author: Ciusyan 5/7/24

package cycle_10_5_7_5_11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
思路重复：
将有序数组转换成二叉搜索树，那么就得知道 BST 的性质：
中序遍历是有序的。
也就是说，一段序列的某一个位置 idx 如果是当前树的 root 节点，那么：
[l, idx) 就是 root.Left 的所有节点。
(idx, r] 就是 root.Right 的所有节点。
所以去递归的构建即可。
*/
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
