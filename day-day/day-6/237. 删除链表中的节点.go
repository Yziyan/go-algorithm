// @Author: Ciusyan 2023/9/11

package day_6

// https://leetcode.cn/problems/delete-node-in-a-linked-list/

func deleteNode2(node *ListNode) {
	// 因为不能拿到前面的节点，所以我们只能从后面的节点入手，
	//	将后一个节点的值覆盖当前需要删除节点的值，
	node.Val = node.Next.Val
	// 	然后删除当前节点的下一个节点即可。
	node.Next = node.Next.Next
}
