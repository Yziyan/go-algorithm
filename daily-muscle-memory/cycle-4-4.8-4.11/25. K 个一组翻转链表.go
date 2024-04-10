// @Author: Ciusyan 4/9/24

package cycle_4_4_8_4_11

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
思路重复：

我们对 K 个一组进行翻转的时候，核心其实就是一组一组的翻转。
先找到要翻转的范围，然后将其翻转的同时，并且要维护整体线条。
比如可以分为：前K -> 中K -> 后K 三个组。
在翻转完 前K 后，将它们的尾部节点记录下来 preTail
然后去翻转 中K 这一组，得到翻转后新的 中头和中尾，然后将 preTail.Next -> 中头。
然后从中尾的 Next 节点开始，去翻转 后K 这个组。

但是我们在对其中在对某一组进行翻转的时候，需要注意：
1.只从 head 翻转到 tail
2.翻转前后，head <-> tail 互换
3.翻转后的 tail.Next -> oldTail.Next
*/
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
