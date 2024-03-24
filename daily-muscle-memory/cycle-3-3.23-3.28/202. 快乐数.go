// @Author: Ciusyan 3/24/24

package cycle_3_3_23_3_28

// https://leetcode.cn/problems/happy-number/description/

/*
*
但是这个题，有一个结论，就是：
如果计算过程中，如果出现了 4，代表就不是开心数，
如果出现了 1，就是开心数
所以可以根据这个题目的一个结论，来反推 code，（类似于死记硬背）
*/
func isHappy(n int) bool {

	for n != 1 && n != 4 {
		// 如果计算过程中，出现了 1 or 4，就说明可以退出了
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}

	return n == 1
}

func isHappy2(n int) bool {

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
