// @Author: Ciusyan 10/28/23

package day_21

import "math"

/**
一块金条切成两半，是需要花费和长度数值一样的铜板
比如长度为20的金条，不管怎么切都要花费20个铜板，一群人想整分整块金条，怎么分最省铜板?
例如，给定数组{10,20,30}，代表一共三个人，整块金条长度为60，金条要分成10，20，30三个部分。
如果先把长度60的金条分成10和50，花费60；再把长度50的金条分成20和30，花费50；一共花费110铜板
但如果先把长度60的金条分成30和30，花费60；再把长度30金条分成10和20，花费30；一共花费90铜板
输入一个数组，返回分割的最小代价
*/

func LessMoneySplitGold(arr []int) int {
	if arr == nil {
		return 0
	}

	// 准备一个最小堆
	heap := NewHeap[int](func(x, y int) int {
		return x - y
	})
	// 将所有 arr 加入最小堆中
	for _, v := range arr {
		heap.Add(v)
	}

	minGold := 0
	// 每一次的代价
	curGold := 0

	// 直到堆里面只有一个元素
	for heap.Size() > 1 {
		// 每次的代价就是取出最小的两个数合并
		curGold = heap.Remove() + heap.Remove()
		minGold += curGold
		// 但是需要将此次代价也加入堆中，用于计算后续的代价
		heap.Add(curGold)
	}

	return minGold
}

// LessMoneySplitGold1 暴力方法
func LessMoneySplitGold1(arr []int) int {
	if arr == nil {
		return 0
	}

	// remainArr 剩余需要合并的数组，preGold 先前的代价 return 合并 remainArr 最小的代价
	var process func(remainArr []int, preGold int) int

	process = func(remainArr []int, preGold int) int {
		if len(remainArr) == 1 {
			// 合并到只剩下一个元素了，就不需要合并了
			return preGold
		}

		var (
			// 最小的代价
			minGold = math.MaxInt
		)

		// 否则枚举所有可能
		l := len(remainArr)
		// 两个 for 就能够列举任意两个数相加了 remainArr[i] + remainArr[j]
		for i := 0; i < l; i++ {
			for j := 1; j < l; j++ {
				curGold := remainArr[i] + remainArr[j]

				// 然后将此时 merge 的两个元素放到新数组中
				mgCurArr := append([]int{}, remainArr...)
				mgCurArr[j] = curGold
				mgCurArr = append(mgCurArr, mgCurArr[i+1:]...)
				nextGold := process(mgCurArr, preGold+curGold)

				if nextGold < minGold {
					minGold = nextGold
				}
			}
		}

		return minGold
	}

	return process(arr, 0)
}
