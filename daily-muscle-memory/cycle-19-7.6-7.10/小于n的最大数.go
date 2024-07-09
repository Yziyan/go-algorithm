// @Author: Ciusyan 2024/7/8

package cycle_19_7_6_7_10

import (
	"math"
	"sort"
)

// getMaxDigitLtD 获取小于指定数字的数字。
func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}

// getMaxNumLtN 获取小于 n 的最大数。
func getMaxNumLTN1(digits []int, n int) int {

	var ndigits []int
	// 获取 n 的每一位数字。
	for n > 0 {
		ndigits = append(ndigits, n%10)
		n /= 10
	}

	// 排序给定的数字数组。
	sort.Ints(digits)

	// 数字写入 map 用于查看是否存在。
	m := make(map[int]struct{})
	for _, v := range digits {
		m[v] = struct{}{}
	}

	// 目标数的每一位数字。
	tdigits := make([]int, len(ndigits))

	// 从高位遍历，尽可能地取相同值，除了最后一位。
	for i := len(ndigits) - 1; i >= 0; i-- {
		if i > 0 {
			// 存在相同数字。优先使用最大的数字
			if _, ok := m[ndigits[i]]; ok {
				tdigits[i] = ndigits[i]
				continue
			}
		}
		// 存在小于当前数字的最大数字。
		if d := getMaxDigitLtD(digits, ndigits[i]); d > 0 {
			tdigits[i] = d
			break
		}

		// 回溯
		for j := i; j < len(ndigits); j++ {
			tdigits[j] = 0
			if d := getMaxDigitLtD(digits, ndigits[j]); d > 0 {
				tdigits[j] = d
				break
			}
			// 最高位也没有小于其的最大数字。
			if j == len(ndigits)-1 {
				tdigits = tdigits[:len(tdigits)-1]
			}
		}
		break
	}

	// 拼接目标数。
	var target int
	for i := len(tdigits) - 1; i >= 0; i-- {
		target *= 10
		if tdigits[i] > 0 {
			target += tdigits[i]
			continue
		}
		target += digits[len(digits)-1]
	}
	return target
}

func getMaxNumLTN(digits []int, n int) int {
	// 将数字转换成数组，方便按位处理
	nums := make([]int, 0, 10)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	// 对 digits 排序，方便取最大值
	sort.Ints(digits)
	dl := len(digits)

	// 将 digits 做一个映射，方便查看是否有相同的数字
	exits := make(map[int]bool, dl)
	for _, di := range digits {
		exits[di] = true
	}

	// 存放每一位可能的结果
	targets := make([]int, len(nums))
	// 对目标数字按位遍历，从高位开始，除去最后一位，前面的每一位尽量使用它本身
BEGIN:
	for i := len(nums) - 1; i >= 0; i-- {
		cur := nums[i]
		if i > 0 {
			// 不是最后一位，才能优先使用相等的数字
			if exits[cur] {
				targets[i] = cur
				continue
			}
		}
		// 说明没有相同的数字，只能看看有没有小于当前位的最大值了
		if ltMax := getLTMax(digits, cur); ltMax != math.MinInt {
			// 说明有比当前位更小的最大值
			targets[i] = ltMax
			// 找到后，直接返回即可
			break
		}

		// 如果来到这里，说明没有比当前位置还小的数字，需要回溯到前一位，重新找一位更小的
		for j := i; j < len(nums); j++ {
			// 都回溯了，先把当前位的结果置为 0，并且只能寻找小于 nums[j] 的最大值
			targets[j] = 0
			if ltMax := getLTMax(digits, nums[j]); ltMax != math.MinInt {
				// 说明找到了，直接 break 掉
				targets[j] = ltMax
				break BEGIN
			}

			if j == len(nums)-1 {
				// 说明到了最后一位都没有某一位有结果，那么说明相同的位数已经没有满足的了，只能减少数字的位数了
				targets = targets[0 : len(targets)-1]
				break BEGIN
			}
		}
	}

	// 现在根据收集的 targets 组装结果
	res := 0
	for i := len(targets) - 1; i >= 0; i-- {
		res *= 10
		if targets[i] > 0 {
			// 说明当时找到了一个合适的值
			res += targets[i]
		} else {
			// 说明没有填写最大值，直接使用 digits 的最大值
			res += digits[dl-1]
		}
	}

	return res
}

// 在有序数组 digits 中，小于 n 的最大值，没有就返回 math.MinInt
func getLTMax(digits []int, n int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < n {
			return digits[i]
		}
	}
	// 说明没有一个数比 n 小
	return math.MinInt
}
