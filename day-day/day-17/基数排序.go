// @Author: Ciusyan 9/26/23

package day_17

// RadixSort 基数排序
func RadixSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	// 找到最大值
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
	}

	// 然后需要对最大值的每一位基数，进行排序，如果是 10 进制，只有 10 位数字，那么我们可以使用计数排序
	for i := 1; i < mx; i *= 10 {
		// i = 1、10、100 依次代表个位数、十位数、百位数...
		countingSort(nums, i)
	}
}

// nums[i] 的基数 = nums[i] / divider % 10
//
//	请注意，一定是对 nums[i] / divider % 10 进行排序
func countingSort(nums []int, divider int) {

	l := len(nums)

	// 默认十进制，所以最多有 10 个桶
	count := make([]int, 10)

	// 将当前基数出现的次数计数
	for i := 0; i < l; i++ {
		// nums[i]/divider%10 代表此次 nums[i] 的基数
		count[nums[i]/divider%10]++
	}

	// 累加 count，count[i] 代表比 i 小的数，有多少个
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 准备一个 help 数组，用于暂存元素应该放在哪个桶
	help := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		// nums[i]/divider%10 代表比它还小的数字出现了多少次，那么它的索引范围就是 0 ~ count[radix]-1
		radix := nums[i] / divider % 10
		count[radix]--
		help[count[radix]] = nums[i] // 注意这里要填写完整的数字，不能只放基数
	}

	// 覆写 nums 数组
	for i := 0; i < l; i++ {
		nums[i] = help[i]
	}
}
