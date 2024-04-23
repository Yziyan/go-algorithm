// @Author: Ciusyan 4/24/24

package cycle_7_4_22_4_26

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var (
		slow = head
		fast = head
	)

	// 找到目标节点的前一个，先让 fast 和 slow 差距 n+1
	for i := 0; i < n+1; i++ {
		if fast == nil {
			// 说明在走的过程中，fast 为 nil 了
			return head.Next
		}
		fast = fast.Next
	}

	// 现在再让 fast 和 slow 同步走
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 当 fast 为 nil 后，其实 slow 就走到了目标节点的前一个节点
	slow.Next = slow.Next.Next

	return head
}
