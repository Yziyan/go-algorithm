// @Author: Ciusyan 2024/7/11

package cycle_20_7_11_7_15

// https://leetcode.cn/problems/sort-colors/

func sortColors(nums []int) {
	l, r := 0, len(nums)-1

	for cur := 0; cur <= r; {
		if nums[cur] < 1 {
			nums[cur], nums[l] = nums[l], nums[cur]
			l++
			cur++
		} else if nums[cur] == 1 {
			cur++
		} else {
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}
	}
}
