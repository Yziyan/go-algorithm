// @Author: Ciusyan 3/24/24

package cycle_3_3_23_3_28

// https://leetcode.cn/problems/happy-number/description/

func isHappy(n int) bool {

	set := make(map[int]bool, 10)
	for n != 1 {
		sum := 0

		for n > 0 {
			// 算出这一次的结果
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		// 然后将 n 换成  sum
		n = sum
		if set[n] {
			// 说明出现循环了
			break
		}
		// 加入 Set
		set[n] = true
	}

	// 看看出没出现 1
	return n == 1
}
