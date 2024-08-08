// @Author: Ciusyan 2024/8/8

package cycle_24_8_05_8_9

import "sort"

// https://leetcode.cn/problems/merge-intervals/description/

/**
思路重复：
1.根据每一个区间对开始位置进行排序。
2.准备两个区间的 l 和 r 代表当前的范围。
3.如果后面的每一个区间开始，如果当前区间的开始位置能和 r 连接上
就将其区间合并了。否则就说明没有链接了，记录一个结果。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l, r, idx := intervals[0][0], intervals[0][1], 0

	for cur := 1; cur < len(intervals); cur++ {
		if intervals[cur][0] > r {
			intervals[idx][0] = l
			intervals[idx][1] = r
			idx++

			l = intervals[cur][0]
			r = intervals[cur][1]
		} else {
			r = max(r, intervals[cur][1])
		}
	}

	return intervals[:idx]
}

func merge2(intervals [][]int) [][]int {

	// 先按照区间的开始位置，对所有区间排序。
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 左边界、右边界、复用索引
	l, r, idx := intervals[0][0], intervals[0][1], 0
	for cur := 0; cur < len(intervals); cur++ {
		if intervals[cur][0] > r {
			// 说明当前遇到的区间，和之前的没有连接了，可以独立开了
			intervals[idx][0] = l
			intervals[idx][1] = r
			idx++

			// 然后更新最新的左右边界
			l = intervals[cur][0]
			r = intervals[cur][1]
		} else {
			// 说明还有连接，看看能否推大右边界
			r = max(r, intervals[cur][1])
		}
	}

	// 最后一组没有推大
	intervals[idx][0] = l
	intervals[idx][1] = r
	idx++

	return intervals[:idx]
}
