// @Author: Ciusyan 2023/7/23

package list

type Queue struct {
	size int64
	head *ListNode
	tail *ListNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int64 {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(v int64) {
	node := NewListNode(v)

	if q.tail == nil {
		// 说明此时队列为空
		q.head = node
	} else {
		// 接在尾部就可以了
		q.tail.next = node
	}

	q.tail = node
	q.size++
}

func (q *Queue) Poll() int64 {
	v := int64(0)
	if q.head != nil {
		v = q.head.value
		// 删除头部节点
		q.head = q.head.next

		q.size--
	}

	if q.head == nil {
		q.tail = nil
	}

	return v
}
