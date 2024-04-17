// @Author: Ciusyan 4/18/24

package cycle_6_4_17_4_21

import "math"

// https://leetcode.cn/problems/reverse-integer/description/

func reverse(x int) int {

	res := 0

	neg := false
	if x < 0 {
		neg = true
		x = -x
	}

	for x != 0 {
		oldRes := res
		lastNum := x % 10
		res = res*10 + lastNum
		if (res-lastNum)/10 != oldRes || res > math.MaxInt32 {
			// 说明中途越界了
			return 0
		}

		x /= 10
	}

	if neg {
		return -res
	}
	return res
}
