// @Author: Ciusyan 2023/7/23

package subject

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	start := head
	// 获取第一组的结尾
	end := kEnd(start, k)
	if end == nil {
		return head
	}

	// 最终返回的节点肯定是第一组的节点
	head = end

	// 翻转
	reverse(start, end)

	// 上一组的末尾
	lastEnd := start

	for lastEnd.Next != nil {
		start = lastEnd.Next
		end = kEnd(start, k)
		if end == nil {
			return head
		}

		reverse(start, end)
		// 更正上一组的末尾节点
		lastEnd.Next = end
		lastEnd = start
	}

	return head
}

// 返回第 n 组的末尾， k 个为一组
func kEnd(head *ListNode, k int) *ListNode {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}

	return head
}

// 对 [start, end) 范围内的节点翻转，并且将这一组反转后的末尾，连接到下一组的开头
func reverse(start, end *ListNode) {
	end = end.Next

	var (
		prev *ListNode
		next *ListNode

		// 原来的第一个节点就是翻转后的最后一个节点
		cur = start
	)

	for cur != end {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// 将翻转后的最后一个节点，连接到下一组的开头
	start.Next = end
}
