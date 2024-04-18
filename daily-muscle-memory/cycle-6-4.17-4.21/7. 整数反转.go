// @Author: Ciusyan 4/18/24

package cycle_6_4_17_4_21

import "math"

// https://leetcode.cn/problems/reverse-integer/description/

/*
*
思路重复：
一种思路是，将 int -> string，对 string 翻转后，再转回 int，
这种思路就不写了，我们写另一种思路：
1.每次取出 x 的最后一位数字 unit，
2.然后变更增加到结果上，res = res*10 + unit

但是在这个过程中，需要对符号、溢出进行判断。方法很多种，
1.按正数计算
2.res 能够转回 preRes，要不然说明越界了
*/
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
