// @Author: Ciusyan 10/4/23

package day_18

// 给定两个可能有环也可能无环的单链表，头节点head1和head2
//	请实现一个函数，如果两个链表相交，请返回相交的第一个节点。如果不相交返回null
//	要求如果两个链表长度之和为N，时间复杂度请达到O(N)，额外空间复杂度请达到O(1)

func GetIntersectNode(headA *ListNode, headB *ListNode) *ListNode {

	loopA := getLoopNode(headA)
	loopB := getLoopNode(headB)

	if loopA == loopB {
		// 这里包含了两种情况：
		//	1. 无环 loopA == nil && loopB == nil
		//	2. 有环 且入环节点相同（说明相交节点在入环节点前）
		return noLoop(headA, headB, loopA)
	}

	// 来到这里有三种情况：
	//	1. A 有环 B 无环
	//  2. B 有环 A 无环
	// 	3. A B 都有环
	//	前两种情况不可能相交，所以我们这里再处理一种即可
	if loopA != nil && loopB != nil {
		// 现在来到这里，A B 都有环，只可能有两种情况：
		// 	1. A B 不相交
		//	2. A B 相交，但是入环节点不同
		// 那么让 cur 围着一个环转一圈，看看能不能碰到另一个环
		cur := loopA.Next
		for cur != loopA {
			if cur == loopB {
				// 来到这里，说明是第二种情况，返回 loopA 和 loopB 都可以
				return loopB
			}
			cur = cur.Next
		}
	}

	// 来到这里，说明统统不可能相交
	return nil
}

// 找到第一个入环的节点，如果没有环，那么返回 nil
func getLoopNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 准备快慢指针
	var (
		slow = head.Next
		fast = head.Next.Next
	)

	// 先让 fast 去追 slow
	for slow != fast {
		if fast == nil || fast.Next == nil {
			// 说明没有环
			return nil
		}

		// 快指针快一步
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 来到这里，快慢指针相遇了，现在让 head 去追 slow
	for slow != head {
		// 一样的速度走
		head = head.Next
		slow = slow.Next
	}

	// 他们一定在第一个入环节点处相遇
	return slow
}

// 链表相交的问题，返回相交节点，没有就返回 nil
// headA 第一条链表
// headB 第二条链表
// end 求到某个节点为基准点
func noLoop(headA *ListNode, headB *ListNode, end *ListNode) *ListNode {

	curA := headA
	curB := headB

	for curA != curB {
		// curA
		if curA != end {
			curA = curA.Next
		} else {
			// 说明 A 走完了，走 B 的路
			curA = headB
		}

		// curB
		if curB != end {
			curB = curB.Next
		} else {
			// 说明 B 走完了，走 A 的路
			curB = headA
		}
	}

	return curA
}
