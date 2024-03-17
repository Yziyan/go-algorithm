// @Author: Ciusyan 3/16/24

package phase_1

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
