// @Author: Ciusyan 2024/8/8

package cycle_24_8_05_8_9

import "sort"

// https://leetcode.cn/problems/merge-intervals/description/

func merge(intervals [][]int) [][]int {

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
