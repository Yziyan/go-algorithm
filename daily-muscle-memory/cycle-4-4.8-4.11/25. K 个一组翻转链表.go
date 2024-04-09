// @Author: Ciusyan 4/9/24

package cycle_4_4_8_4_11

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		dummyHead = &ListNode{Next: head} // 虚拟头结点
		preTail   = dummyHead             // 每一组翻转时，上一组翻转后的结尾
	)

	for head != nil {
		// 从上一组的尾部节点开始
		tail := preTail
		// 找到 k 个节点的结尾
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 如果某一次为 nil 了，说明没有 k 个节点了，直接返回
				return dummyHead.Next
			}
		}
		// 到这里，这 K 个一组的节点已就绪：(head -> ... -> tail -> tail.Next)
		// 将其翻转，变成 (tail -> ... -> head -> tail.Next)
		head, tail = reverseEnd(head, tail)
		// 将上一组与这一组链接起来
		preTail.Next = head
		// 下一组的开始
		head = tail.Next
		// 对下一组链接做准备，这一组翻转后的 tail，就是下一组的 preTail
		preTail = tail
	}

	return dummyHead.Next
}

// 从 start 开始翻转，到 end 结束，返回翻转后，新的头部
// 3 -> 4 -> 5 -> 6 -> nil， start = 3, end = 5
// 5 -> 4 -> 3 -> 6， head = 5, tail = 3
func reverseEnd(start, end *ListNode) (head, tail *ListNode) {
	prev := end.Next
	// 一会需要角色互换，头变尾，尾变头
	head = end
	tail = start

	// 只翻转到 end
	for prev != end {
		next := start.Next
		start.Next = prev
		prev = start
		start = next
	}

	// 翻转完成后，返回互换后的头和尾
	return
}
