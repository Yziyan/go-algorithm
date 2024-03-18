// @Author: Ciusyan 3/15/24

package cycle_1_3_13_3_17

// https://leetcode.cn/problems/path-sum-iii/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var process func(root *TreeNode, targetSum int) int
	// 从 root 出发，能有多少种解法
	process = func(root *TreeNode, targetSum int) int {
		if root == nil {
			// 说明没有节点
			return 0
		}
		res := 0
		val := root.Val
		if val == targetSum {
			// 说明到目前为止，已经有一种方案可以找出路径和了
			res++
			// 但是还得继续往下寻找，因为节点的值可能有负数
		}
		// 去左右子树进行搜索
		res += process(root.Left, targetSum-val)  // 往左树找 targetSum-val
		res += process(root.Right, targetSum-val) // 往右树找 targetSum-val

		return res
	}

	// 先计算从根节点出发，能有的结果数
	res := process(root, targetSum)
	// 但是还需要看看，左子树、右子树中有多少满足条件的解法
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

/*
	思路重复 / 伪代码

思路分析：
想要求出所有满足条件的路径和数量，数量其实就等同于求解出：

	根树的路径总和数量 + 左子树的路径总和的数量 + 右子树的路径总和的数量

左右子树的路径总和数量都好解，对于主函数递归调用即可。

那么根树的路径总和怎么求呢？所以就得设计递归函数了。

思路重复：
既然没路过一个节点，它的路径就会相应的增加，那么我们将其还需要凑的路径总和传入进去计算
即可设计这样的递归函数： process(root, remainSum) int
其含义代表：从 root 出发，剩余 remainSum 要凑，返回能凑出 remainSum 的方法数。

如果 root 是空树，可以理解成没法凑。
否则取出当前的 val，如果刚好能凑出 remainSum，就代表是一种结果
但是不能返回了，还需要继续看看，左右子树能凑出多少种情况，将他们相加才是结果。
去计算左右子树的时候，那么 remainSum 就需要减去当前 val 的值了。
*/
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 根 + 左 + 右
	res := process2(root, targetSum)
	res += pathSum2(root.Left, targetSum)  // 左子树出发的结果
	res += pathSum2(root.Right, targetSum) // 右子树出发的结果

	return res
}

// 从 root 出发，凑出 remainSum 的数量
func process2(root *TreeNode, remainSum int) int {
	if root == nil {
		return 0
	}

	res := 0
	val := root.Val
	if val == remainSum {
		// 说明已经得到一种结果了
		res++
	}

	// 不管之前有没有得到结果，都需要继续看看左右子树能凑出 remainSum-val 的数量
	res += process2(root.Left, remainSum-val)
	res += process2(root.Right, remainSum-val)

	return res
}
