// @Author: Ciusyan 10/30/23

package day_21

// https://leetcode.cn/problems/ipo/description/

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	l := len(profits)
	// 将所有 profits 和 capital 关联起来 project[i][0] 花费、project[i][1] 利润
	project := make([][]int, l)
	for i := 0; i < l; i++ {
		project[i] = []int{capital[i], profits[i]}
	}

	// 准备两个堆
	// 1.按照花费最小（最小堆）
	minCapital := NewHeap[[]int](func(x, y []int) int {
		return x[0] - y[0]
	})
	// 2.按照利润最高（最大堆）
	maxProfit := NewHeap[[]int](func(x, y []int) int {
		return y[1] - x[1]
	})

	for i := 0; i < l; i++ {
		minCapital.Add(project[i])
	}

	// 收集 k 个项目
	for i := 0; i < k; i++ {
		for minCapital.Size() != 0 && minCapital.Get()[0] <= w {
			// 来到这里，说明有基金可以做项目
			maxProfit.Add(minCapital.Remove())
		}

		// 如果没有项目可做，直接返回了
		if maxProfit.Size() == 0 {
			return w
		}

		// 有的话，就把最挣钱的项目做了
		w += maxProfit.Remove()[1]
	}

	return w
}
