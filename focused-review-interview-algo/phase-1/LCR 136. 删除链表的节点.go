// @Author: Ciusyan 3/16/24

package phase_1

func deleteNode136(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		// 说明删除第一个节点
		head = head.Next
	}

	// 否则先找到前驱节点，然后删除对应的节点
	prev := head
	cur := head.Next

	for cur != nil {
		if cur.Val == val {
			// 说明找到要删除的节点了，直接删除
			prev.Next = cur.Next
			break
		}

		// 否则继续往下遍历
		prev = cur
		cur = cur.Next
	}

	return head
}
