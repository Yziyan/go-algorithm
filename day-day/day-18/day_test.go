// @Author: Ciusyan 10/2/23

package day_18

import (
	"fmt"
	"math/rand"
	"testing"
)

func NewNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func randList() *ListNode {
	dummyHead := NewNode(0)
	tail := dummyHead

	l := rand.Intn(15)
	for i := 0; i < l; i++ {
		node := NewNode(rand.Intn(9))
		tail.Next = node
		tail = node
	}

	return dummyHead.Next
}

func printList(head *ListNode) {
	var str []byte
	for head != nil {
		str = fmt.Appendf(str, "%d->", head.Val)
		head = head.Next
	}
	l := len(str) - 2
	if l <= 0 {
		fmt.Println("nil")
		return
	}
	fmt.Println(string(str[:l]))
}

func TestListPartition(t *testing.T) {
	head := randList()
	printList(head)
	pivot := rand.Intn(5)
	fmt.Printf("pivot = %d\n", pivot)
	head = ListPartition(head, pivot)
	printList(head)
}
