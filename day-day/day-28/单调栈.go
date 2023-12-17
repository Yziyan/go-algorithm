// @Author: Ciusyan 12/17/23

package day_28

// 数组元素不可重复，返回值含义：[2]int 代表：[i位置左边第一个比我小的数， i位置右边第一个比我小的数]
// arr: [2, 1, 3]
// res: [
//
//	       0: [-1, 1]
//			  1: [-1, -1]
//			  2: [1, -1]
//	     ]
func getNearLessNoRepeat(arr []int) [][2]int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	n := len(arr)
	res := make([][2]int, n)
	// 准备一个单调栈，存放的是索引
	stack := NewStack()

	for i := range arr {
		// 将 arr[i] 添加进入栈，但是要保证单调性（栈底->栈顶，按从小到大排列）
		for stack.Size() != 0 && arr[i] < arr[stack.Peek()] {
			// 说明当前元素，比栈顶元素还要小，需要先腾出栈顶
			// 下面的操作都是堆索引而言，先弹出栈顶
			popIdx := stack.Pop()
			// 左边第一个比 popIdx 大的数，默认弹出后左边没数了
			leftLessIdx := -1
			if stack.Size() != 0 {
				// 说明栈还有数，
				leftLessIdx = stack.Peek()
			}
			// 存储结果，谁弹出，就记录谁
			res[popIdx][0] = leftLessIdx
			res[popIdx][1] = i
		}
		// 来到这里，说明栈顶干净了，可以放心加入了
		stack.Push(i)
	}

	// 当遍历完后，需要清空栈里的元素
	for stack.Size() != 0 {
		popIdx := stack.Pop()
		// 也是弹出就记录结果
		leftLessIdx := -1
		if stack.Size() != 0 {
			// 说明弹出后还有元素
			leftLessIdx = stack.Peek()
		}
		// 记录结果
		res[popIdx][0] = leftLessIdx
		res[popIdx][1] = -1 // 因为是自然弹出的，没有人比它还小，能让他弹出
	}

	return res
}
