// @Author: Ciusyan 5/18/24

package cycle_12_5_17_5_21

// https://leetcode.cn/problems/plus-one/

/**
思路重复：
从最后一位开始加一，如果某一位是 9，加一会进位，将这一位变为 0，然后去下一位相加。
如果不是 9，直接加一后，就可以返回了。

如果全部数字都遍历完了还没有返回，说明全部都是 9，需要多一位。
那么就返回 len(digits) + 1 的数组，并且将第一位设置为 1，其余全为 0
*/

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			// 来到这里，说明加完没有进位了，可以终止了
			return digits
		}

		// 能到这里，说明要去加下一位
		digits[i] = 0
	}

	// 来到这里，之前说明上面没返回，全进位了
	digits = make([]int, len(digits)+1)
	digits[0] = 1

	return digits
}
