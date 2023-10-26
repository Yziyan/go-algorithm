// @Author: Ciusyan 10/26/23

package day_21

import "sort"

// 一些项目要占用一个会议室宣讲，会议室不能同时容纳两个项目的宣讲，给你每一个项目开始的时间和结束的时间。
// 你来安排宣讲的日程，要求会议室进行的宣讲的场次最多，返回最多的宣讲场次

// BestArrange 使用的贪心策略：按照宣讲结束时间最早
// arranges：所有的宣讲场次 arranges[i][0] 和 arranges[i][1] 分别代表第 i 场宣讲的开始时间和结束时间
func BestArrange(arranges [][2]int) int {
	if arranges == nil || len(arranges) == 0 {
		return 0
	}

	// 先按照会议的结束时间排序
	sort.Slice(arranges, func(i, j int) bool {
		return arranges[i][1] < arranges[j][1]
	})

	// 宣讲结束的时间线
	timeLine := 0
	res := 0
	// 遍历每一场会议
	for i := 0; i < len(arranges); i++ {
		// 如果当前宣讲的时间比当前会议的开始时间还小，就说明不能安排这一场宣讲
		if arranges[i][0] < timeLine {
			continue
		}

		// 这里说明可以安排当前场次
		res++
		// 记得把时间线调整为当前场结束的时间
		timeLine = arranges[i][1]
	}

	return res
}

// BestArrange1 暴力方法
func BestArrange1(arranges [][2]int) int {
	if arranges == nil || len(arranges) == 0 {
		return 0
	}

	// 定义一个递归函数
	// remainArranges：剩余的宣讲安排，done：已经安排的场数，timeLine：上一场会议的结束时间
	// 返回能够安排的最大宣讲场数
	var process func(remainArranges [][2]int, done int, timeLine int) int
	process = func(remainArranges [][2]int, done int, timeLine int) int {
		// 如果没有会议了，就返回之前已经安排过的会议
		if len(remainArranges) == 0 {
			return done
		}

		// 说明还有会议可以安排
		// 默认最大可以安排 done 场
		maxCount := done
		// 对所有剩余的会议挨个尝试
		for i, cur := range remainArranges {
			// 如果当前宣讲的开始时间比上一场结束时间还早，说明不能安排
			if cur[0] < timeLine {
				continue
			}

			// 来到这里说明可以安排
			// 将当前场次删除
			rmCurArranges := append([][2]int{}, remainArranges[0:i]...)
			rmCurArranges = append(rmCurArranges, remainArranges[i+1:]...)

			// 然后递归的去处理，
			nextMaxCount := process(rmCurArranges, done+1, cur[1])
			if nextMaxCount > maxCount {
				// 收到后面场次安排的结果后，来看看是不是比当前这个结果大，是的话就安排上
				maxCount = nextMaxCount
			}
		}

		return maxCount
	}

	// 默认所有场次可安排，安排了 0 场，上一场结束时间是 0
	return process(arranges, 0, 0)
}
