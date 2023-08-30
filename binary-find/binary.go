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
