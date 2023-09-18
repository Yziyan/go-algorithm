// @Author: Ciusyan 9/18/23

package day_12

// NationalFlag 将 nums 中使用 nums[end] 作为分隔条件，将 [begin, end] 分隔成三段
//
//	小于 nums[end] 的放在左边，等于的放在中间，大于的放在后面。最终返回中间的开始和结尾位置的索引
func NationalFlag(nums []int, begin, end int) []int {
	if begin > end {
		// 越界了
		return []int{-1, -1}
	}
	if begin == end {
		// 只有一个数
		return []int{begin, end}
	}

	// 准备几个指针
	var (
		l   = begin
		cur = begin
		r   = end - 1
	)

	// 只要还有没有看过的数，就得扫
	for cur <= r {
		if nums[cur] < nums[end] {
			// 小于当前值，交换 l 和 cur，并且都往后走
			nums[cur], nums[l] = nums[l], nums[cur]
			cur++
			l++
		} else if nums[cur] == nums[end] {
			// 直接往下走一个就可以
			cur++
		} else {
			// 大于当前值，交换 r 和 cur，r 往前走，cur 不变，因为后面交换过来的值还没看过
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}
	}
	// 退出来后，一定排好序了，但是我们还需要将 nums[end] 归位，与 cur 交换就可以了
	nums[cur], nums[end] = nums[end], nums[cur]

	// 返回的边界是 l 和 cur
	return []int{l, cur}
}
