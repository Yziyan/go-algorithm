// @Author: Ciusyan 9/24/23

package day_15

import "sort"

// CoverMax 线段最大重合问题：给定一组线段，最大的重合数
//
//	比如 lines = [ [1, 4], [4, 9], [2, 8], [3, 10] -> 3
func CoverMax(lines [][]int) int {
	if lines == nil || len(lines) < 1 {
		return 0
	}

	// 先按照线段的起始时间升序排列
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})

	// 准备一个最小堆，用于排序结尾位置
	heap := NewHeap()

	res := 0
	// 遍历所有线段
	for i := 0; i < len(lines); i++ {
		// 每次用当前段的起始位置与堆中最近的结束位置比较
		for !heap.IsEmpty() && lines[i][0] >= heap.Get() {
			// 说明堆里有还能够贯穿的线段，并且与当前线段最近的结束位置都贯穿不了，说明已经不会贯穿之后了
			heap.Remove()
		}
		// 然后将当前段的结束位置加入堆中
		heap.Add(lines[i][1])
		res = max(res, heap.Size())
	}

	return res
}
