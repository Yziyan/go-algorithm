// @Author: Ciusyan 1/20/24

package day_31

// https://leetcode.cn/problems/sqrtx/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 3 {
		return 1
	}

	// 否则使用二分去找结果
	var (
		res = 0

		// 在 [1, x] 上二分
		begin = 1
		end   = x
		mid   = 0
	)

	// 一定要分到最后
	for begin <= end {
		// begin + end / 2
		mid = begin + ((end - begin) >> 1)
		// 看看 mid 的平方，和 x 谁大
		powMid := mid * mid
		if powMid == x {
			// 如果都相等了，直接返回啊，说明能开整的
			return mid
		}

		if powMid < x {
			// 说明是一个可能的结果，先记录一个结果
			res = mid
			// 但是右边可能还有更合适的，继续往右走
			begin = mid + 1
		} else {
			// 说明选大了，得往左二分
			end = mid - 1
		}
	}

	return res
}
