// @Author: Ciusyan 9/26/23

package day_17

// CountingSort 计数排序
func CountingSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	l := len(nums)
	// 先找出最大最小值
	var (
		mx = nums[0]
		mn = nums[0]
	)
	for i := 1; i < l; i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
		if nums[i] < mn {
			mn = nums[i]
		}
	}

	// 准备一一个 Count 数组，长度就是 Max - Min + 1
	count := make([]int, mx-mn+1)
	// 挨个计数填充 Count 数组
	for i := 0; i < l; i++ {
		// nums[i] - 最小值，那么能保证数组最小索引至少是 1，支持负数
		count[nums[i]-mn]++
	}

	// 改造 Count 数组，变成累加和
	//	 此刻的 count 数组，count[i] 代表：比 i 小的数有几个
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	// 辅助数组，用于承接计数后的顺序
	help := make([]int, l)
	// 从后往前填写
	for i := l - 1; i >= 0; i-- {
		// 比 nums[i]-mn 还小的数有 3 个，那么 nums[i]-mn 的索引就是 0~2，那么最右边的位置就是 2 位置
		count[nums[i]-mn]--
		help[count[nums[i]-mn]] = nums[i]
	}

	// 将排序后的数组覆盖掉原始数组
	for i := 0; i < l; i++ {
		nums[i] = help[i]
	}
}
