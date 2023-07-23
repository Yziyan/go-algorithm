// @Author: Ciusyan 2023/7/23

package list

func NewDoubleListNode(v int64) *DoubleListNode {
	return &DoubleListNode{V: v}
}

type DoubleListNode struct {
	V    int64
	prev *DoubleListNode
	next *DoubleListNode
}

type Deque struct {
	Size int
	Head *DoubleListNode
	Tail *DoubleListNode
}

func (d *Deque) LPush(v int64) {
	node := NewDoubleListNode(v)
	if d.Tail == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.next = d.Head
		d.Head.prev = node
		d.Head = node
	}

	d.Size++
}

func (d *Deque) RPush(v int64) {
	node := NewDoubleListNode(v)
	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.prev = d.Tail
		d.Tail.next = node
		d.Tail = node
	}
	d.Size++
}

func (d *Deque) LPoll() int64 {
	if d.Head == nil {
		return 0
	}

	d.Size--
	res := d.Head.V
	d.Head = d.Head.next
	if d.Head != nil {
		d.Head.prev = nil
	} else {
		d.Tail = nil
	}

	return res
}

func (d *Deque) RPoll() int64 {
	if d.Tail == nil {
		return 0
	}

	d.Size--
	res := d.Tail.V
	d.Tail = d.Tail.prev
	if d.Tail != nil {
		d.Tail.next = nil
	} else {
		d.Head = nil
	}

	return res
}
