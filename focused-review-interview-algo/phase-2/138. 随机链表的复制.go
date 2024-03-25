// @Author: Ciusyan 3/25/24

package phase_2

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 源链表：1 -> 2 -> 3 -> nil
	// 新链表：1 -> 1` -> 2 -> 2` -> 3 -> 3` -> nil
	cur := head
	for cur != nil {
		next := cur.Next
		// 拷贝节点
		cur.Next = &Node{
			Val:  cur.Val,
			Next: next,
		}

		// 然后去下一个位置接
		cur = next
	}

	newHead := head.Next
	// 然后填充拷贝链表的 Random 节点
	cur = head
	for cur != nil {
		// 取出 cp 节点
		cp := cur.Next
		if cur.Random != nil {
			// 说明这个拷贝节点需要接 Random
			cp.Random = cur.Random.Next
		}
		// 然后去下一个原始节点
		cur = cp.Next
	}

	// 分离两条链表
	cur = head
	for cur != nil {
		// 取出 cp 节点
		cp := cur.Next
		// 源链表更新了
		cur.Next = cp.Next
		cur = cp.Next
		if cur != nil {
			// 说明后面还有节点
			cp.Next = cur.Next
		}

	}

	return newHead
}
