// @Author: Ciusyan 2023/8/31

package day_1

import (
	"fmt"
	"math/rand"
	"testing"
)

func randList(l int) *ListNode {

	head := &ListNode{}
	tail := head

	for i := 1; i <= l; i++ {
		node := &ListNode{Val: rand.Intn(i)}
		tail.Next = node
		tail = node
	}

	return head.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	head := randList(8)
	printList(head)
	//head = ReverseList(head)
	printList(head)

	head = RemoveListValue(head, 0)
	printList(head)
}
