// @Author: Ciusyan 2023/7/27

package dfs

// https://leetcode.cn/problems/permutations/

func Permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	l := len(nums)
	track := make([]int, l)
	used := make([]bool, l)
	results := make([][]int, 0)
	dfs(0, &results, nums, track, used)

	return results
}

func dfs(level int, results *[][]int, nums, track []int, used []bool) {
	if level == len(nums) {
		// 说明到了最底层，记录结果
		*results = append(*results, append([]int(nil), track...))

		return
	}

	// 到达这里，利用所有可能构建结果
	for i, v := range nums {
		if used[i] {
			continue
		}

		// 来到这里可以选择
		track[level] = v
		used[i] = true
		// 往下一层钻
		dfs(level+1, results, nums, track, used)
		// 还原现场
		used[i] = false
	}
}
