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
