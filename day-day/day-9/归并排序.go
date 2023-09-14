// @Author: Ciusyan 9/14/23

package day_9

// MergeSort1 归并排序
//
//	递归版实现
func MergeSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	sort(arr, 0, len(arr))
}

// 对 arr 中 [begin, end) 排序
func sort(arr []int, begin, end int) {
	// 如果只有一个元素了，就没必要排序了
	if end-begin < 2 {
		return
	}

	// 求出中点位置
	mid := begin + (end-begin)>>1
	// 对左右分别进行归并排序
	sort(arr, begin, mid)
	sort(arr, mid, end)
	// 排序完成后，对左右两边进行合并
	merge(arr, begin, mid, end)
}

func merge(arr []int, begin, mid, end int) {
	// 准备一个大数组
	tempArr := make([]int, end-begin)
	i := 0
	// 准备左右指针
	p1 := begin
	p2 := mid

	for p1 < mid && p2 < end {
		// 比较，谁小放谁，默认 p1 位置的小
		temp := arr[p1]
		p1++

		if arr[p2] < temp {
			// 说明当初错付了，选择对的，别忘了还回旧的
			temp = arr[p2]
			p1--
			p2++
		}

		// 小的给到 i 位置
		tempArr[i] = temp
		i++
	}
	// 来到这里，要么 p1 越界了，要么 p2 越界了
	for p1 < mid {
		tempArr[i] = arr[p1]
		p1++
		i++
	}

	for p2 < end {
		tempArr[i] = arr[p2]
		p2++
		i++
	}

	// 将旧数组改写
	for _, v := range tempArr {
		arr[begin] = v
		begin++
	}
}

// MergeSort2 非递归实现
func MergeSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 步长
	mergeSize := 1
	l := len(arr)

	// 只要 mergeSize 小于数组长度，就可以尝试
	for mergeSize < l {
		// 每一次从第一组开始，只要组不超过数组长度就可以去合并
		begin := 0
		for begin < l {
			if mergeSize >= l-begin {
				// 要是左边都灭有一个步长的长度，就没必要合并了
				break
			}
			// 找到第一组开头左边界
			mid := begin + mergeSize
			// 找到右边的结束（mid + 步长 or 剩余长度）
			end := mid + min(mergeSize, l-mid)

			// 合并 [begin, mid) 和  [mid, end)
			merge(arr, begin, mid, end)

			begin = end
		}

		// 在做完一次操作后，步长扩大 2
		mergeSize <<= 1
	}

}
