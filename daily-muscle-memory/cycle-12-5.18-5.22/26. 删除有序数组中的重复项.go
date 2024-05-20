// @Author: Ciusyan 5/21/24

package cycle_12_5_18_5_22

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {

	i := 1
	for cur := 1; cur < len(nums); cur++ {
		if nums[cur] == nums[i-1] {
			// 说明这个元素现在填过了
			continue
		}
		// 说明第一次出现
		nums[i] = nums[cur]
		i++
	}

	return i
}
