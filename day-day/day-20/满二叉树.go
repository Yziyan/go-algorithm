// @Author: Ciusyan 10/18/23

package day_20

// 一棵树是否是满二叉树取决于：一颗高为 h，有 n 个节点的二叉树满足 2^h - 1 = n

// IsFull 方法一：求解出 root 的 h 和 n
func IsFull(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type info struct {
		height int
		size   int
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			return &info{}
		}

		// 收集左右子树信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			// 高度和 size 都默认是 left 的
			height = leftInfo.height
			size   = leftInfo.size
		)

		// 尝试使用右子树推大
		if rightInfo.height > height {
			height = rightInfo.height
		}

		if rightInfo.size > size {
			size = rightInfo.size
		}

		return &info{
			// 都别忘了要加自己
			height: height + 1,
			size:   size + 1,
		}
	}

	rootInfo := getInfo(root)
	// 判断是不是呢？可以看看高度和节点数量的关系
	return (1<<rootInfo.size)-1 == rootInfo.size
}

// IsFull2 方法二：直接判断左右子树是否是满二叉树，并且的高度是否一样
func IsFull2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type info struct {
		height int
		isFull bool
	}

	var getInfo func(root *TreeNode) *info
	getInfo = func(root *TreeNode) *info {
		if root == nil {
			// 空树默认高度是 0，是 Full 的
			return &info{
				isFull: true,
			}
		}

		// 收集左右子树信息
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			// 高度默认是 left 的
			height = leftInfo.height
			isFull = false
		)

		// 必须左右子树都是 Full 并且 高度要一致
		isFull = leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height

		// 尝试使用右子树推大
		if rightInfo.height > height {
			height = rightInfo.height
		}

		return &info{
			// 都别忘了要加自己
			height: height + 1,
			isFull: isFull,
		}
	}

	// 判断是不是呢？可以看看高度和节点数量的关系
	return getInfo(root).isFull
}
