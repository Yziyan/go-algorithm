// @Author: Ciusyan 10/3/23

package day_18

// https://leetcode.cn/problems/copy-list-with-random-pointer/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 遍历原链表，先将新节点建起来挂载在旧节点的后面
	//	1 -> 2 -> 4 -> nil
	//	1 -> 1` -> 2 -> 2` -> 4 -> 4` -> nil

	cur := head
	for cur != nil {
		next := cur.Next
		// 新建节点，并且串联起来
		cur.Next = &Node{
			Val:  cur.Val,
			Next: next,
		}

		cur = next
	}

	newHead := head.Next
	// 来到这里，节点已经串起来了，从头开始遍历，为 copyList：
	//	1. 接上 random
	cur = head
	for cur != nil {
		// 此次的  cp 节点
		cp := cur.Next
		// 为 cp 链接上 random
		if cur.Random != nil {
			cp.Random = cur.Random.Next
		}
		// 跳到下一个旧节点即可
		cur = cp.Next
	}

	//  2. 修正 next
	cur = head
	for cur != nil {
		// 此次的  cp 节点
		cp := cur.Next
		// 修正原链表的 Next
		cur.Next = cp.Next
		// 修正新链表的 Next
		if cp.Next != nil {
			cp.Next = cp.Next.Next
		}
		// 跳到下一个旧节点即可
		cur = cur.Next
	}

	return newHead
}
