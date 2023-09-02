// @Author: Ciusyan 2023/9/2

package day_3

// https://leetcode.cn/problems/reverse-nodes-in-k-group/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 先将开头保存
	start := head
	// 先找到第一组，将其翻转后，得到新的头结点
	end := kEnd(start, k)
	if end == nil {
		// 说明不足 K 个节点，不用翻转了
		return head
	}

	// 来到这里，先将 head 保存为第一组的末尾，作为新链表的头结点
	head = end

	// 对第一组进行翻转
	reverse(start, end)

	// 翻转后，start 就变成了第一组的末尾节点
	lastEnd := start

	// 如果上一组翻转后还链接这下一组的开头，说明还需要继续翻转
	for lastEnd.Next != nil {
		// 现在进行新一组的翻转
		start = lastEnd.Next
		end = kEnd(start, k)
		if end == nil {
			return head
		}

		// 来到这里说明还有 k 个，可以翻转
		reverse(start, end)
		// 但是在翻转后，需要更正上一组的结尾的指向
		lastEnd.Next = end
		// 现在上一组已经变成新的一组了
		lastEnd = start
	}

	return head
}

// 返回某组的第 k 个节点，刚好是第 k 个
func kEnd(start *ListNode, k int) *ListNode {
	for k > 1 && start != nil {
		start = start.Next
		k--
	}

	return start
}

// 对 [start, end) 的节点进行翻转，但是翻转后需要将 start 链接到下一组的头部
func reverse(start, end *ListNode) {
	// 因为要翻转到 end，所以需要先将 end 后移
	end = end.Next

	var (
		prev *ListNode
		next *ListNode

		// 需要将 start 保留
		oldStart = start
	)

	// 标准翻转代码，但是只翻转到 end
	for start != end {
		next = start.Next
		start.Next = prev
		prev = start
		start = next
	}

	// 翻转结束后，需要将老的开头接入下一组
	oldStart.Next = end
}
