// @Author: Ciusyan 2023/8/10

package arrary

// 选择排序：
//
//	在待排序序列中选择最小值或者最大值，将其放到序列的开头或者结尾
func changeSort(arr []int) {
	if arr == nil {
		return
	}

	for begin := 0; begin < len(arr)-1; begin++ {

		minIdx := begin
		for end := begin + 1; end < len(arr); end++ {
			// 找到最小的一个索引
			if arr[minIdx] > arr[end] {
				minIdx = end
			}
		}

		// 交换他们的值
		arr[begin], arr[minIdx] = arr[minIdx], arr[begin]
	}
}
