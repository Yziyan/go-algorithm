// @Author: Ciusyan 10/14/23

package day_20

func StackSort(num []int) {
	if len(num) == 0 {
		return
	}

	// 当做 Stack
	stackA := Stack(num)

	// 准备一个新的栈，准备给它排序
	stackB := NewStack()
	stackATop := 0
	// 回弹的数量
	backSize := 0

	for stackA.Size() != 0 {
		// 每轮会处理一个 A 栈的元素
		stackATop = stackA.Pop()

		if stackB.Size() == 0 || stackB.Top() <= stackATop {
			// 如果另一个栈都没有元素，那就加进去即可
			stackB.Push(stackATop)
			continue
		}

		// 来到这里说明要大一些，现在不适合放在 B 栈，需要找出位置
		// 即将弹回 stackA 中，并记录弹出的数量
		for stackB.Size() != 0 && stackB.Top() > stackATop {
			stackA.Push(stackB.Pop())
			backSize++
		}

		// 来到这里，说明可以存放了
		stackB.Push(stackATop)

		// 然后再弹回来
		for backSize > 0 {
			stackB.Push(stackA.Pop())
			backSize--
		}
	}

	// 现在 B 是有序的那个
	for i := len(num) - 1; i >= 0; i-- {
		num[i] = stackB.Pop()
	}
}
