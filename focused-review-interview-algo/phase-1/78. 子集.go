// @Author: Ciusyan 3/16/24

package phase_1

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	var dfs func(level int, nums []int, track []int, res *[][]int)
	dfs = func(level int, nums []int, track []int, res *[][]int) {
		if level == len(nums) {
			// 说明搜到了最后一层，记录一个结果
			*res = append(*res, append([]int{}, track...))
			return
		}

		// 否则尝试所有可能的情况
		//	1.不选择当前元素
		dfs(level+1, nums, track, res)
		// 	2.选择当前元素
		track = append(track, nums[level])
		// 然后往下一层搜索
		dfs(level+1, nums, track, res)
		// 搜索回来后，记得还原现场
		track = track[:len(track)-1]
	}

	var (
		res   = make([][]int, 0, len(nums))
		track = make([]int, 0, len(nums))
	)

	// 从第一层开始去搜索
	dfs(0, nums, track, &res)

	return res
}
