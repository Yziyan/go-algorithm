// @Author: Ciusyan 2023/7/27

package arrary

func mergeKSortedArray(arrays [][]int) []int {
	if len(arrays) == 1 {
		return arrays[0]
	}

	mid := len(arrays) >> 1
	// 拆分成子问题
	left := mergeKSortedArray(arrays[:mid])
	right := mergeKSortedArray(arrays[mid:])

	// 合并左右数组，
	return mergeSortedArray(left, right)
}

func mergeSortedArray(arr1, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	res := make([]int, 0, l1+l2)

	i1, i2 := 0, 0
	// 必须小于其中一个切片的长度
	for i1 < l1 && i2 < l2 {
		if arr1[i1] < arr2[i2] {
			res = append(res, arr1[i1])
			i1++
		} else {
			res = append(res, arr2[i2])
			i2++
		}
	}

	// 但是需要将余下的一个数组添加到结果上
	res = append(res, arr1[i1:]...)
	res = append(res, arr2[i2:]...)

	return res
}
