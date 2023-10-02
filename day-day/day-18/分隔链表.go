// @Author: Ciusyan 10/2/23

package day_18

// ListPartition 给定一个单链表的头节点head，给定一个整数n，将链表按n划分成左边<n、中间==n、右边>n
func ListPartition(head *ListNode, pivot int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 准备三条链表
	var (
		// 小于 pivot
		lH = &ListNode{}
		lT = lH

		// 等于 pivot
		mH = &ListNode{}
		mT = mH

		// 大于 pivot
		rH = &ListNode{}
		rT = rH
	)

	for head != nil {
		next := head.Next
		head.Next = nil
		// 判断当前节点，应该接哪一条链表
		if head.Val < pivot {
			lT.Next = head
			lT = head
		} else if head.Val == pivot {
			mT.Next = head
			mT = head
		} else {
			rT.Next = head
			rT = head
		}

		head = next
	}

	// 来到这里，说明链表已经被分割成三条；还需要将三条链表串起来
	// 将中间和右边串起来
	mT.Next = rH.Next
	// 将左边和中间串起来
	lT.Next = mH.Next

	return lH.Next
}
