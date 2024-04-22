// @Author: Ciusyan 4/22/24

package cycle_7_4_22_4_26

// https://leetcode.cn/problems/permutations-ii/description/

func permuteUnique(nums []int) [][]int {

	isRepeat := func(nums []int, level, willSwap int) bool {
		for i := level; i < willSwap; i++ {
			if nums[i] == nums[willSwap] {
				return true
			}
		}

		return false
	}

	var dfs func(nums []int, level int, res *[][]int)
	dfs = func(nums []int, level int, res *[][]int) {
		if level == len(nums) {
			*res = append(*res, append([]int{}, nums...))
			return
		}

		for cur := level; cur < len(nums); cur++ {
			if isRepeat(nums, level, cur) {
				// 说明 [level ... cur) 有过 nums[cur] 这个值了，
				// 就别去交换了，因为他们最终会有重复的交集
				continue
			}
			// 说明可以交换
			nums[cur], nums[level] = nums[level], nums[cur]
			dfs(nums, level+1, res)
			// 记得还原现场
			nums[cur], nums[level] = nums[level], nums[cur]
		}
	}

	res := make([][]int, 0, 100)
	dfs(nums, 0, &res)

	return res
}
