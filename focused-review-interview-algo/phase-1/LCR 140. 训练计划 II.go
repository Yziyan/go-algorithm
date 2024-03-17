// @Author: Ciusyan 3/16/24

package phase_1

func trainingPlan(head *ListNode, cnt int) *ListNode {
	if head == nil {
		return head
	}

	var (
		fast = head
		slow = head
	)

	// 先让 fast 先走 cnt 步
	for cnt > 0 && fast != nil {
		fast = fast.Next
		cnt--
	}

	// 然后 fast 和 slow 现在同步走
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
