// @Author: Ciusyan 2023/8/30

package bit_operation

import "fmt"

// Swap 交换 arr 中，i 和 j 位置的值
//
//	使用位运算，要求 i 和 j 不能相等，相等会有问题
func Swap(arr []int, i, j int) {
	if i == j {
		return
	}

	// 在 Go 语言当然可以这样写，很方便。相当于一种使用临时变量的写法
	// arr[i], arr[j] = arr[j], arr[i]

	// 使用位运算呢？
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

// OddTimesNum1 arr 中，只有一种数，出现奇数次
func OddTimesNum1(arr []int) int {
	if arr == nil {
		return -1
	}

	// 只需要将所有数都 ^ 一遍，其他的数字都是偶数次，总能交换为 0，
	// 剩余一个奇数次的数字 ^ 0，还是本身，即找出了对应的数字
	res := 0
	for i := range arr {
		res ^= arr[i]
	}

	return res
}

// OddTimesNum2 arr 中，只有两个数，出现了奇数次，返回他们的和
func OddTimesNum2(arr []int) int {
	if arr == nil {
		return -1
	}

	// eor = 奇数个数 ^ 另一个奇数个数
	eor := 0
	for _, v := range arr {
		eor ^= v
	}

	// 取出 eor 二进制的最后一个 1，=> eor & (-eor)
	eorLast1 := eor & (^eor + 1)
	// 10 10 011 100 01 01 110 100
	// 2, 2, 3,  4,  1, 1,  6,  4
	// eor = 3 ^ 6
	// eor` = 001
	// 3 ^ 1 ^ 1

	// nums1 = 其中一个出现奇数次的数字
	num1 := 0
	for _, v := range arr {
		if (v & eorLast1) != 0 {
			num1 ^= v
		}
	}

	// nums2 = 奇数个数 ^ 另一个奇数个数 ^ 其中一个奇数次的数字 = 另一个数字
	nums2 := eor ^ num1
	fmt.Println(num1, nums2)

	return num1 + nums2
}
