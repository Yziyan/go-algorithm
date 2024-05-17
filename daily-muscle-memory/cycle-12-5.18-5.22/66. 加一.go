// @Author: Ciusyan 5/18/24

package cycle_12_5_18_5_22

// https://leetcode.cn/problems/plus-one/

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
