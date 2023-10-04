// @Author: Ciusyan 10/2/23

package day_18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestGetIntersectNode(t *testing.T) {

	// 1->2->3->4->5->6->7->null
	head1 := NewNode(1)
	head1.Next = NewNode(2)
	head1.Next.Next = NewNode(3)
	head1.Next.Next.Next = NewNode(4)
	head1.Next.Next.Next.Next = NewNode(5)
	head1.Next.Next.Next.Next.Next = NewNode(6)
	head1.Next.Next.Next.Next.Next.Next = NewNode(7)

	// 0->9->8->6->7->null
	head2 := NewNode(0)
	head2.Next = NewNode(9)
	head2.Next.Next = NewNode(8)
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next // 8->6

	// 1->2->3->4->5->6->7->4...
	head3 := NewNode(1)
	head3.Next = NewNode(2)
	head3.Next.Next = NewNode(3)
	head3.Next.Next.Next = NewNode(4)
	head3.Next.Next.Next.Next = NewNode(5)
	head3.Next.Next.Next.Next.Next = NewNode(6)
	head3.Next.Next.Next.Next.Next.Next = NewNode(7)
	head3.Next.Next.Next.Next.Next.Next.Next = head3.Next.Next.Next // 7->4

	// 0->9->8->2...
	head4 := NewNode(0)
	head4.Next = NewNode(9)
	head4.Next.Next = NewNode(8)
	head4.Next.Next.Next = head3.Next // 8->2

	// 0->9->8->6->4->5->6..
	head5 := NewNode(0)
	head5.Next = NewNode(9)
	head5.Next.Next = NewNode(8)
	head5.Next.Next.Next = head3.Next.Next.Next.Next.Next // 8->6

	testCase := []struct {
		name string

		// 参数
		headA *ListNode
		headB *ListNode

		// 返回值
		wantNode *ListNode
	}{
		{
			name: "no loop",
			// 1->2->3->4->5->[6]->7->null
			headA: head1,
			// 0->9->8->[6]->7->null
			headB: head2,
			// 6
			wantNode: head2.Next.Next.Next,
		},
		{
			name: "loop 环前相交",
			// 1->[2]->3->[loop4]->5->6->7->4...
			headA: head3,
			// 0->9->8->[2]->3->[loop4]->5->6->7->4...
			headB: head4,
			// 2
			wantNode: head3.Next,
		},
		{
			name: "loop 入环节点不同",
			// 1->2->3->[loop4]->5->6->7->4...
			headA: head3,
			// 0->9->8->[loop6]->7->4->5->6...
			headB: head5,
			// 6 or 4 都是可以的
			wantNode: head3.Next.Next.Next.Next.Next,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			got := GetIntersectNode(tc.headA, tc.headB)
			// 这里是比较它们的地址必须是一个
			assert.True(t, got == tc.wantNode)
		})
	}

}
