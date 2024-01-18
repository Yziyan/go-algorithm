// @Author: Ciusyan 1/18/24

package day_31

import "sort"

// https://leetcode.cn/problems/merge-intervals/

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return intervals
	}

	// 先按照区间的开始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var (
		begin = intervals[0][0] // 区间的开始位置
		end   = intervals[0][1] // 区间的结束位置
		idx   = 0               // 用于复用 intervals 索引
	)

	for cur := 0; cur < len(intervals); cur++ {
		// 当前为 cur 号区间
		if intervals[cur][0] > end {
			// 说明开始位置比之前最晚的结束时间大，需要记录结果了
			intervals[idx][0] = begin
			intervals[idx][1] = end
			idx++
			// 但是要记得更新 begin 和 end
			begin = intervals[cur][0]
			end = intervals[cur][1]
		} else {
			// 说明还没有比之前 end 晚的，重试更新 end
			if intervals[cur][1] > end {
				end = intervals[cur][1]
			}
		}
	}

	// 但是我们是在每一次区间被后来者冲的时候，才去更新的结果，最后一组肯定没人能够冲它了
	// 需要手动设置最后一组的结果
	intervals[idx][0] = begin
	intervals[idx][1] = end
	idx++

	// [0, idx) 是有效的结果区间
	return intervals[:idx]
}
