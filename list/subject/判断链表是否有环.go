// @Author: Ciusyan 2023/7/26

package subject

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var (
		slow = head
		fast = head.Next
	)

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
