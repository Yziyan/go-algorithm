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
