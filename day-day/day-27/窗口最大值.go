// @Author: Ciusyan 12/12/23

package day_27

/**
窗口内最大值或最小值更新结构的实现
假设一个固定大小为 W 的窗口，依次划过 arr，
返回每一次滑出状况的最大值
例如，arr = [4,3,5,4,3,3,6,7], W = 3
返回：[5,5,5,4,6,7]
*/

func slidingWindowMaxArray(arr []int, w int) []int {
	if arr == nil || len(arr) < w {
		return nil
	}

	n := len(arr)

	// 先构造结果，要滑完所有窗口，肯定有 n-w+1 个结果
	res := make([]int, n-w+1)
	// 准备一个双端队列，用于：
	// 快速搜集窗口的最大值，队列里存储索引，严格按照队头 -> 队尾，从小到大的顺序
	queue := NewDoubleQueue()
	// 收集结果使用的索引
	idx := 0

	// 窗口的左边界，从 0 - n
	for r := 0; r < n; r++ {
		// 每进来一个，就尝试从右边入队
		for queue.Size() != 0 && arr[r] >= arr[queue.GetRight()] {
			// 说明队列有元素，并且新划入窗口的元素比队尾的要大，
			// 说明队尾存在这个窗口，已经无意义了，毕竟永远不可能比新滑入窗口的要大。直接弹出不要了
			queue.PollRight()
		}

		// 说明可以加入了，记得是添加索引
		queue.OfferRight(r)

		// 加入过后，还要看窗口左侧有没有使队头中的索引失效
		if r-w == queue.GetLeft() {
			// 说明队头已经不能囊括在窗口中了，即使它再大，也需要将其抛弃，
			queue.PollLeft()
		}

		// 看看有没有形成一个完整的窗口，形成了就收集结果
		if r-w+1 >= 0 {
			// 说明肯定形成一个窗口了
			// 队头的就是此时窗口的最大值
			res[idx] = arr[queue.GetLeft()]
			idx++
		}
	}

	return res
}
