// @Author: Ciusyan 9/19/23

package day_13

import (
	"math/rand"
)

// https://leetcode.cn/problems/sort-an-array/description/

func sortArray(nums []int) []int {
	if nums == nil {
		return nil
	}

	// 对 nums 进行快速排序
	quickSort(nums, 0, len(nums))

	return nums
}

// 对 nums 的 [begin, end) 进行快速排序
func quickSort(nums []int, begin, end int) {
	if end-begin < 2 {
		return
	}
	//[0, 1) * (end - begin) -> [0, end-begin) + begin -> [begin, end)
	r := begin + int(float64(end-begin)*rand.Float64())
	// 随机一些，随机选取一个座位轴点
	nums[begin], nums[r] = nums[r], nums[begin]

	// 获取轴点元素
	pivotL, pivotR := getPivotPoint(nums, begin, end)
	// 对轴点左右也进行快速排序
	quickSort(nums, begin, pivotL)
	quickSort(nums, pivotR+1, end)
}

// 利用 L 位置作为轴点，对其进行荷兰国旗问题排序，返回的是，两个分隔点
//
//	比排序后是 [1, 4, 2, 3, 5, 5, 5, 7, 6, 9, 8]
//		5 是基准点，所以返回值是 [4, 6]
func getPivotPoint(nums []int, begin, end int) (int, int) {
	// 越界了
	if begin >= end {
		return -1, -1
	}
	// 只有一个元素
	if end-begin < 2 {
		return begin, end
	}

	// 因为外界传入的 end，是不可取的
	end--
	// 准备几个指针
	var (
		l   = begin + 1
		cur = l
		r   = end
	)

	// 每一个数都需要查看
	for cur <= r {
		// 与 Nums[begin] 作为基准
		if nums[cur] < nums[begin] {
			// 比当前元素小，与 l 交换，并且 cur 和 l 都往后走
			nums[cur], nums[l] = nums[l], nums[cur]
			cur++
			l++
		} else if nums[cur] == nums[begin] {
			// 相等，只有 cur 往后走
			cur++
		} else {
			// 大于当前元素，与 r 交换，但是只有 r 往前走
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		}
	}

	// 需要将 轴点 归位
	nums[l-1], nums[begin] = nums[begin], nums[l-1]

	// 最终返回值，就是 l 和 cur
	return l - 1, r
}
