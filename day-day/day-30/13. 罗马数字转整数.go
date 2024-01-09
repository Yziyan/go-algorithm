// @Author: Ciusyan 1/9/24

package day_30

// https://leetcode.cn/problems/roman-to-integer/

func romanToInt(s string) int {
	l := len(s)
	// 先将每一个字符映射成数字
	nums := make([]int, l)
	for i, c := range s {
		num := 0
		switch c {
		case 'I':
			num = 1
		case 'V':
			num = 5
		case 'X':
			num = 10
		case 'L':
			num = 50
		case 'C':
			num = 100
		case 'D':
			num = 500
		case 'M':
			num = 1000
		}

		nums[i] = num
	}

	// 默认用最后一位开始相加，因为个位后面不可能有数了，就不用判断了
	res := nums[l-1]
	// 遍历相加每一个数字
	for i := 0; i < l-1; i++ {
		if nums[i+1] > nums[i] {
			// 说明要加上当前的相反数
			nums[i] = -nums[i]
		}

		res += nums[i]
	}

	return res
}
