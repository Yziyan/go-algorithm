// @Author: Ciusyan 10/19/23

package day_20

// MaxSubBSTTree 最大 BST 子树
func MaxSubBSTTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// info 信息
	type info struct {
		maxBstSubTreeSize int // 所求结果
		size              int // 节点数量
		max               int // 最大值
		min               int // 最小值
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			return nil
		}

		// 收集左右子树的信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			// 设置默认值
			maxBstSubTreeSize = 0
			size              = 1
			mx                = root.Val
			mn                = root.Val
		)

		// 先尝试使用左右子树的信息更新信息
		if leftInfo != nil {

			if leftInfo.max > mx {
				mx = leftInfo.max
			}

			if leftInfo.min < mn {
				mn = leftInfo.min
			}

			size += leftInfo.size
		}

		if rightInfo != nil {

			if rightInfo.max > mx {
				mx = rightInfo.max
			}

			if rightInfo.min < mn {
				mn = rightInfo.min
			}

			size += rightInfo.size
		}

		// 可能的结果有三种
		// 1、在左树
		if leftInfo != nil {
			maxBstSubTreeSize = leftInfo.maxBstSubTreeSize
		}
		// 2、在右树
		if rightInfo != nil {
			// 与现有的做个比较
			if rightInfo.maxBstSubTreeSize > maxBstSubTreeSize {
				maxBstSubTreeSize = rightInfo.maxBstSubTreeSize
			}
		}

		// 3、在整棵树
		// 所以需要判断，是否是 BST，并且 Size 有多少

		var (
			leftIsBst  = false
			rightIsBst = false
			isBst      = false
		)

		// 左树为 nil 或者满足 maxBstSubTreeSize = size
		if leftInfo == nil || leftInfo.maxBstSubTreeSize == leftInfo.size {
			leftIsBst = true
		}
		// 右树也满足
		if rightInfo == nil || rightInfo.maxBstSubTreeSize == rightInfo.size {
			rightIsBst = true
		}

		if leftIsBst && rightIsBst {
			// 只有当左右子树都是 BST 的前提，才有必要看 BST
			leftMaxLess := leftInfo == nil || leftInfo.max < root.Val
			rightMinThan := rightInfo == nil || rightInfo.min > root.Val
			// 只有左子树的最大值小于 root、右子树的最小值大于 root，同时成立才满足
			if leftMaxLess && rightMinThan {
				isBst = true
			}
		}

		if isBst {
			// 说明 root 是 BST，那么看整棵树的 Size 多大
			if size > maxBstSubTreeSize {
				maxBstSubTreeSize = size
			}
		}

		return &info{
			maxBstSubTreeSize: maxBstSubTreeSize,
			size:              size,
			max:               mx,
			min:               mn,
		}
	}

	return getInfo(root).maxBstSubTreeSize
}
