// @Author: Ciusyan 10/17/23

package day_20

// https://leetcode.cn/problems/diameter-of-binary-tree/description/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Info 信息
	type Info struct {
		// 最大直径
		maxDiameter int
		// 高度
		height int
	}

	// 获取信息的函数
	var _getInfo func(root *TreeNode) *Info
	_getInfo = func(root *TreeNode) *Info {
		if root == nil {
			return &Info{}
		}

		// 收集左右子树的信息
		leftInfo := _getInfo(root.Left)
		rightInfo := _getInfo(root.Right)

		// 利用左右子树的信息构建出 Info
		var (
			// 先默认是左子树的直径和高度
			maxDiameter = leftInfo.maxDiameter
			height      = leftInfo.height
		)

		// 查看右子树能否比左边高
		if rightInfo.height > height {
			height = rightInfo.height
		}

		// 右子树的最大直径可能比自己左边大
		if rightInfo.maxDiameter > maxDiameter {
			maxDiameter = rightInfo.maxDiameter
		}

		// 有可能左边和右边能够加起来(穿过 root)
		throughRoot := leftInfo.height + rightInfo.height
		if throughRoot > maxDiameter {
			maxDiameter = throughRoot
		}

		return &Info{
			maxDiameter: maxDiameter,
			height:      height + 1, // 还得把自己的高度算进去
		}
	}

	// 调用 _getInfo
	return _getInfo(root).maxDiameter
}
