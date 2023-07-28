// @Author: Ciusyan 2023/7/28

package dfs

// https://leetcode.cn/problems/subsets/submissions/

func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	var (
		track []int
		res   [][]int
	)

	// 从第 0 层开始搜索
	subsetsDfs(0, nums, track, &res)

	return res
}

func subsetsDfs(level int, nums, track []int, res *[][]int) {
	if level == len(nums) {
		// 记录一个结果
		*res = append(*res, append([]int{}, track...))

		return
	}

	// 列出所有可能的选择
	//	1、选择当前数字
	track = append(track, nums[level])
	subsetsDfs(level+1, nums, track, res)

	// 	2、不选择当前数字（需要先还原现场，删除最后一个数字）
	track = track[0 : len(track)-1]
	subsetsDfs(level+1, nums, track, res)
}
