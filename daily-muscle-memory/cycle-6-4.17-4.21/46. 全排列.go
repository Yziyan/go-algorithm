// @Author: Ciusyan 4/21/24

package cycle_6_4_17_4_21

// https://leetcode.cn/problems/permutations/

/**
思路重复：
对于全排列，说到排列组合，必不可想的就是：dfs。
这个题不会有重复元素，我们所有元素，只需要保证都只使用一次去排即可，即我们在 dfs 的过程中，
最好能够记录 xx 元素的使用状态。在搜索完成从下层回来的时候，再还原现场，去进行下一个元素的搜索。

但是我们可以有一个更好的解法就是：我们每一次直接使用 nums 作为轨迹，每一次去搜索之前，我们都已经确认了 [0 ... level） 的值
即我们只需要去搜索确定 [level ... n）的值。即可。
所以我们在 level 去探索的过程中，挨个与 level 位置互换。从下层搜索结束回来后再换回来，即可得到所有排列组合的结果。
*/

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
