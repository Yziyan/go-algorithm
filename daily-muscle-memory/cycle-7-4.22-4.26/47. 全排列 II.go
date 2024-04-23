// @Author: Ciusyan 4/22/24

package cycle_7_4_22_4_26

// https://leetcode.cn/problems/permutations-ii/description/

/**
思路重复：
当我们会写昨天的全排列后，这个进阶版，应该怎么做呢？我们核心就是要去掉重复的情况。因为给定的元素有可能会重复。
那么我们在交换前，我们就需要将其需要交换的位置，和前面的 level 位置的元素比较。如果已经有 willSwap 的值了，就别换了
比如 level = 1，nums = [1, 2, 3, 4, 3, 5]，willSwap = 4
那么这个值，是不能换掉的，因为换掉了，之后也会重复。

*/

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
