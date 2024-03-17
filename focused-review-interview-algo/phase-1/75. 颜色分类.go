// @Author: Ciusyan 3/17/24

package phase_1

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		pivot = 1 // 轴点值

		l   = 0
		cur = 0
		r   = len(nums) - 1
	)

	// 直至确定完所有元素
	for cur <= r {
		if nums[cur] < pivot {
			// 比轴点小，要放轴点左边，
			nums[cur], nums[l] = nums[l], nums[cur]
			// 换过来后，相当于位置确定了
			l++
			// 并且换过来的元素也检查过了
			cur++
		} else if nums[cur] == pivot {
			// 和轴点相等，待定
			cur++
		} else {
			// 说明比轴点大，应该放在它后面，先把 r 位置的换过来再说
			nums[cur], nums[r] = nums[r], nums[cur]
			//换过来后，末尾的值就已经确定位置了，但是要检查当前被换过来的值
			r--
		}
	}
}
