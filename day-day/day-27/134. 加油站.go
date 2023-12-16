// @Author: Ciusyan 12/16/23

package day_27

// https://leetcode.cn/problems/gas-station/description/

func canCompleteCircuit(gas []int, cost []int) int {
	status := getStatus(gas, cost)
	if status == nil {
		return -1
	}

	for i, v := range status {
		if v {
			// 只需要找到第一个能跑通的加油站即可
			return i
		}
	}

	// 说明都不能跑通
	return -1
}

// 返回所有位置，能否跑通的情况
func getStatus(gas, cost []int) []bool {
	if gas == nil || cost == nil || len(gas) == 0 || len(gas) != len(cost) {
		return nil
	}

	n := len(gas)
	m := n << 1
	// 先整合改造数组
	status := make([]int, m)
	for i := 0; i < n; i++ {
		status[i] = gas[i] - cost[i]
		status[i+n] = gas[i] - cost[i]
	}

	// 累加起来
	for i := 1; i < m; i++ {
		status[i] += status[i-1]
	}

	// 准备一个滑动窗口，维护着从 i 位置出发，最薄弱的位置
	queueMin := NewDoubleQueue()
	// 先填充第一个窗口
	for r := 0; r < n; r++ {
		for queueMin.Size() != 0 && status[r] <= status[queueMin.GetRight()] {
			// 说明当前要进入窗口的值，比窗口末尾的值还小，加入后会违反从小到大的顺序
			queueMin.PollRight()
		}
		// 这里可以放心加入了，但是别加错了，是索引
		queueMin.OfferRight(r)
	}

	var (
		// 窗口的前缀，用于求解 i 位置，原始的累加和
		prefix = 0

		// 目前窗口的左边界，用于检查窗口维护的值是否失效
		l = 0
		// 窗口的右边界
		r = n

		res = make([]bool, n)
	)

	// 不能越界
	for r < m {
		if status[queueMin.GetLeft()]-prefix >= 0 {
			// 说明最薄弱的位置，原始的累加和不会被攻陷，可以跑通环路
			res[l] = true
		}

		// 需要将窗口整体往后滑动一个单位
		for queueMin.Size() != 0 && status[r] <= status[queueMin.GetRight()] {
			// 说明当前值加入窗口会违反窗口从小到大的顺序
			queueMin.PollRight()
		}
		queueMin.OfferRight(r)

		// 还需要校验窗口维护的最小值是否会失效
		if queueMin.GetLeft() == l {
			// 说明滑动后就失效了
			queueMin.PollLeft()
		}

		// 更新前缀、滑动窗口边界
		prefix = status[l]
		l++
		r++
	}

	return res
}
