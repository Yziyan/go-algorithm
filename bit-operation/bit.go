// @Author: Ciusyan 2023/8/30

package bit_operation

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
