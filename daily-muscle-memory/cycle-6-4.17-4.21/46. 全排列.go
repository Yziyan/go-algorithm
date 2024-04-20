// @Author: Ciusyan 4/21/24

package cycle_6_4_17_4_21

// https://leetcode.cn/problems/permutations/

func permute(nums []int) [][]int {

	var dfs func(level int, nums []int, res *[][]int)
	dfs = func(level int, nums []int, res *[][]int) {
		if level == len(nums) {
			// 到达最后一层，保存搜集的结果，再返回
			*res = append(*res, append([]int{}, nums...))
			return
		}

		// 从 level 开始，列举所有可能，目前，[0...level) 的值，都已经确定了，接下来都是为了确定 level 的值
		for cur := level; cur < len(nums); cur++ {
			// 交换 level 和 cur 的位置，代表确定 level 的值
			nums[cur], nums[level] = nums[level], nums[cur]
			// 去下一层搜索
			dfs(level+1, nums, res)
			// 但是搜索完成后，记得还原现场
			nums[cur], nums[level] = nums[level], nums[cur]
		}
	}

	res := make([][]int, 0, 100)
	// 从第零层开始搜索
	dfs(0, nums, &res)

	return res
}
