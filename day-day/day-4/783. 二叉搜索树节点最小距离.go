// @Author: Ciusyan 2023/9/3

package day_4

import "math"

// https://leetcode.cn/problems/minimum-distance-between-bst-nodes/

// 二叉树的递归套路
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}

	// 得到左右子树的最小差值
	leftMin := minDiffInBST(root.Left)
	rightMin := minDiffInBST(root.Right)

	// 取出左右子树的最小差值
	minRes := leftMin
	if rightMin < minRes {
		minRes = rightMin
	}

	// 还需要与前驱节点和后继节点比较
	// 获取前驱节点与当前节点的差值
	if root.Left != nil {
		prec := root.Left
		// 一直查找到左子树的最右边
		for prec.Right != nil {
			prec = prec.Right
		}
		precDiff := root.Val - prec.Val
		if precDiff < minRes {
			minRes = precDiff
		}
	}

	// 获取后继节点与当前节点的差值
	if root.Right != nil {
		succ := root.Right
		// 一直找到右子树的最左边
		for succ.Left != nil {
			succ = succ.Left
		}
		succDiff := succ.Val - root.Val
		if succDiff < minRes {
			minRes = succDiff
		}
	}

	return minRes
}

// 使用中序遍历求解
func minDiffInBST1(root *TreeNode) int {
	// 最好重置一下变量的初始值
	prev = nil
	minDiff = math.MaxInt
	// 利用中序遍历是有序的求解
	inorder(root)

	// 中序遍历过后，填充完毕了，直接返回即可
	if minDiff == math.MaxInt {
		return 0
	}

	return minDiff
}

// 中序遍历求解
var (
	prev    *TreeNode
	minDiff = math.MaxInt
)

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	// 具体的操作
	if prev != nil {
		// 尝试更新最小值
		diff := root.Val - prev.Val
		if diff < minDiff {
			minDiff = diff
		}
	}

	prev = root

	inorder(root.Right)
}
