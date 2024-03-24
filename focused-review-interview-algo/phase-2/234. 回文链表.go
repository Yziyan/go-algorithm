// @Author: Ciusyan 3/24/24

package phase_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var (
		// 先通过快慢指针，找到终点的前一个节点
		slow = head
		fast = head.Next.Next
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 然后从终点开始，对链表进行翻转，翻转到终点
	tailHead := reverse(slow.Next)

	cur := head
	tail := tailHead
	res := true
	for cur != nil {
		if cur.Val != tail.Val {
			res = false
		}

		cur = cur.Next
		tail = tail.Next
	}

	// 还原回去
	slow.Next = reverse(tailHead)

	return res
}

func reverse(head *ListNode) *ListNode {
	var (
		prev *ListNode
	)
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
