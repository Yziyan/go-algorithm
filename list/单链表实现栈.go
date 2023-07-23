// @Author: Ciusyan 2023/7/23

package list

type Stack struct {
	size int64
	head *ListNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Size() int64 {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(v int64) {
	node := NewListNode(v)
	if s.head == nil {
		s.head = node
	} else {
		// 使用头插法
		node.next = s.head
		s.head = node
	}

	s.size++
}

func (s *Stack) Pop() int64 {
	res := int64(0)

	if s.head != nil {
		res = s.head.value
		s.head = s.head.next

		s.size--
	}

	return res
}
