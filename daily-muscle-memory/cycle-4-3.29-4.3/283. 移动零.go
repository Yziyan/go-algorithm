// @Author: Ciusyan 3/29/24

package cycle_4_3_29_4_3

// https://leetcode.cn/problems/move-zeroes/

/**
思路重复：
最简的一种方式，其实就是两层 for 循环，挨个从后往前扫描，
当扫描到是 0 的时候，就将此位置至目前后面第一个零，全部往前挪动。
但是勒，这样就比较耗时了。

更好的做法是，使用双指针，逆向扫描：
准备一个 cur，和 first，分别代表：当前的位置 和 目前第一个零所在点。first <= cur。
那么从前往后扫描：
1.遇到 0，直接 cur++ 后跳过即可
2.遇到非 0，看看当前位置是否和 first 位置重合，如果重合 cur 和 first 都一起 ++，
否则将 cur 位置的元素，放在 first 位置后再一起 ++

*/

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	for cur, first := 0, 0; cur < len(nums); cur++ {
		if nums[cur] == 0 {
			// 说明遇到零了，直接跳过
			continue
		}
		// 来到这里，说明不是 0
		if cur > first {
			// 说明前面有 0，需要交换
			nums[first] = nums[cur]
			nums[cur] = 0
		}
		first++
	}

}

func moveZeroes222(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		n   = len(nums)
		cur = 0
		fz  = 0
	)

	for cur < n {
		if nums[cur] == 0 {
			// 可以跳过，因为有 fz 在前面守着
			cur++
			continue
		}

		// 说明不是 0，得看看 fz 落后没有，如果 fz 落后了，说明前面有 0，需要换位置
		if cur != fz {
			// 说明需要讲当前这个不为 0 的数往前挪到 fz 的位置
			nums[fz] = nums[cur]
			nums[cur] = 0
		}
		cur++
		fz++
	}
}

func moveZeroes2(nums []int) {
	if len(nums) < 2 {
		return
	}

	var (
		// 准备双指针
		right = len(nums) - 1
	)

	for cur := right - 1; cur >= 0; cur-- {
		if nums[cur] != 0 {
			continue
		}
		// 说明等于 0，需要移动
		for begin := cur; begin < right; begin++ {
			nums[begin] = nums[begin+1]
		}
		nums[right] = 0
		right--
	}
}
