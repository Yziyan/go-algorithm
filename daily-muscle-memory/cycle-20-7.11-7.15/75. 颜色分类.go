// @Author: Ciusyan 2024/7/11

package cycle_20_7_11_7_15

// https://leetcode.cn/problems/sort-colors/

/**
思路重复：
经典的荷兰国旗问题。
准备三指针：l, r, cur
用 cur 从左往右扫描，
* 遇到比 1 小的，就和 l 交换，并且 cur 和 l 都加一
* 遇到和 1 相等的，cur 往后加一
* 遇到比 1 大的，就和 r 交换，并且 r 往前减一
*/

func sortColors1(nums []int) {
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
