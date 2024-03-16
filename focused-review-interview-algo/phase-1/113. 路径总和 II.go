// @Author: Ciusyan 3/16/24

package phase_1

func pathSum(root *TreeNode, targetSum int) [][]int {

	// 从 root 开始，一直搜索到根节点，找出满足 target 路径
	var process func(root *TreeNode, target int, track *[]int, res *[][]int)
	process = func(root *TreeNode, target int, track *[]int, res *[][]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			// 说明到叶子节点了，
			if root.Val != target {
				// 说明这条路径和不为 target
				return
			}
			// 说明需要收集一条合法路径
			curTrack := append([]int(nil), *track...)
			*res = append(*res, append(curTrack, root.Val))
			return
		}

		// 否则记录当前轨迹后去左右子树搜索
		*track = append(*track, root.Val)
		process(root.Left, target-root.Val, track, res)
		process(root.Right, target-root.Val, track, res)
		// 但是再回来之前，记得还原现场，删除
		*track = (*track)[:len(*track)-1]
	}

	var (
		track = make([]int, 0, 1)
		res   = make([][]int, 0, 1)
	)

	// 去进行 dfs
	process(root, targetSum, &track, &res)

	return res
}
