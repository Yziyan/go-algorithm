// @Author: Ciusyan 5/21/24

package cycle_12_5_17_5_21

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

/**
思路重复：
删除有序数组中重复的项，我们从前往后，准备俩指针，扫描过去即可。
具体怎么坐呢？
i、cur 俩指针，
当 nums[cur] == nums[i-1] 的时候，说明这个数字已经取出过了，直接去下一个位置。
去下个位置做同样的操作即可，
直至所有位置都检查完毕。
 */

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
