// @Author: Ciusyan 10/4/23

package day_18

// https://leetcode.cn/problems/intersection-of-two-linked-lists/

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	curA := headA
	curB := headB

	// 直接比较内存地址
	for curA != curB {
		// curA
		if curA != nil {
			curA = curA.Next
		} else {
			// 如果 curA 先走完了，那么就接着 B 走
			curA = headB
		}

		// curB
		if curB != nil {
			curB = curB.Next
		} else {
			// 如果 curB 先走完了，那么就接着 A 走
			curB = headA
		}

	}

	// 走到这里，要么返回的是相交节点，要么返回的是 nil 代表不相交
	return curA
}
