// @Author: Ciusyan 3/23/24

package cycle_3_3_23_3_28

// https://leetcode.cn/problems/find-the-sum-of-encrypted-integers/

func sumOfEncryptedInt(nums []int) int {

	res := 0
	for _, num := range nums {
		// num 的最大位是 mx，有多少位1
		mx, base := 0, 0
		for num > 0 {
			// 有几位，到时候就有多少位
			mx = max(num%10, mx)
			// 每进来一次，就要在末尾多增加一个 1，有多少位，就有多少个 1
			base = base*10 + 1
			// 去算下一位，去掉刚计算的最后一位
			num /= 10
		}
		// 此次 num 的产出就是：位数个1*最大位
		res += base * mx
	}
	return res
}
