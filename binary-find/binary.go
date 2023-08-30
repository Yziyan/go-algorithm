// @Author: Ciusyan 2023/8/29

package binary_find

// Exsit 查看 num 是否存在于 sortedArr 中
func Exsit(sortedArr []int, num int) bool {
	if sortedArr == nil {
		return false
	}

	var (
		begin = 0
		end   = len(sortedArr)
		mid   int
	)

	// 对 [begin, end) 进行二分
	for begin < end {
		mid = begin + (end-begin)>>1

		if sortedArr[mid] == num {
			return true
		}

		if sortedArr[mid] > num {
			// 说明在前面
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return false
}

// NearestIndex 在 arr 中，找满足 >= num 最左边的位置
func NearestIndex(arr []int, num int) int {
	if arr == nil {
		return -1
	}

	var (
		begin = 0
		end   = len(arr)
		mid   int
		index = -1
	)

	// [begin, end)
	for begin < end {
		mid = begin + (end-begin)>>1

		// 等于的时候也往左走，必须要二分到底
		if arr[mid] >= num {
			// 需要记录一下找到的比 num 大的数
			index = mid

			end = mid
		} else {
			begin = mid + 1
		}
	}

	return index
}

// LessIndex arr 不一定是有序的
func LessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	l := len(arr)

	// 处理左边界
	if l == 1 || arr[0] < arr[1] {
		return 0
	}

	// 处理右边界
	if arr[l-1] < arr[l-2] {
		return l - 1
	}

	var (
		begin = 1
		end   = l - 2
		mid   int
	)

	for begin < end {
		mid = begin + (end-begin)>>1

		// 如果中间比左边大，那就往左边靠
		// 如果中间比右边大，那就往右边靠
		// 如果比左右两边都小，那么就说明这是一个局部最小值了嘛
		if arr[mid] > arr[mid-1] {
			end = mid
		} else if arr[mid] > arr[mid+1] {
			begin = mid + 1
		} else {
			return mid
		}
	}

	return begin
}
