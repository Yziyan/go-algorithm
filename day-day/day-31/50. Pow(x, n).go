// @Author: Ciusyan 1/18/24

package day_31

import "math"

// https://leetcode.cn/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	// 准备一个 pow，变成 x^pow 次幂
	pow := n
	if n == math.MinInt {
		// 如果是系统最小值，将其加一，要不然取相反数后就溢出了
		pow = math.MinInt + 1
	}
	if pow < 0 {
		// 如果 pow 是负数，计算它的正次幂
		// pow = -pow
		pow = ^pow + 1
	}

	// 利用快速幂的方式求解
	res := float64(1)
	t := x // x^1

	for pow != 0 {
		// 说明还没计算完成
		if (pow & 1) == 1 {
			// 说明需要将其乘到结果里去
			res *= t
		}

		// 不管怎样，t 都需要自乘
		t *= t
		pow >>= 1
	}

	if n == math.MinInt {
		// 如果之前是系统最小值，那么少算了一个 x
		res *= x
	}

	if n < 0 {
		// 说明之前是负数次幂
		return 1 / res
	}

	return res
}
