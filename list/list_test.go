// @Author: Ciusyan 2023/7/22

package list_test

import (
	"github.com/Yziyan/go-algorithm/list"
	"testing"
)

func TestStack(t *testing.T) {
	stack := list.NewStack()

	// Test Size and IsEmpty on an empty stack
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected stack size to be 0, but got %d", size)
	}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty, but it's not")
	}

	// Test Push and Size on a non-empty stack
	stack.Push(10)
	stack.Push(20)

	if size := stack.Size(); size != 2 {
		t.Errorf("Expected stack size to be 2, but got %d", size)
	}

	if stack.IsEmpty() {
		t.Error("Expected stack not to be empty, but it is")
	}

	// Test Pop on a non-empty stack
	if v := stack.Pop(); v != 20 {
		t.Errorf("Expected popped value to be 20, but got %d", v)
	}

	if size := stack.Size(); size != 1 {
		t.Errorf("Expected stack size to be 1 after pop, but got %d", size)
	}

	// Test Pop on an empty stack
	stack.Pop() // Pop the last element

	if v := stack.Pop(); v != 0 {
		t.Errorf("Expected popped value to be 0 on an empty stack, but got %d", v)
	}

	if size := stack.Size(); size != 0 {
		t.Errorf("Expected stack size to be 0 after popping the last element, but got %d", size)
	}

	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty after popping the last element, but it's not")
	}
}

func TestQueue(t *testing.T) {
	// 创建一个空队列
	q := list.NewQueue()
	// 测试队列的大小和空状态
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, got %d", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, got false")
	}
	// 向队列中添加三个元素
	q.Push(1)
	q.Push(2)
	q.Push(3)
	// 测试队列的大小和空状态
	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", q.Size())
	}
	if q.IsEmpty() {
		t.Errorf("Expected queue to be not empty, got true")
	}
	// 测试队列的出队操作
	if q.Poll() != 1 {
		t.Errorf("Expected queue poll to return 1, got %d", q.Poll())
	}
	if q.Poll() != 2 {
		t.Errorf("Expected queue poll to return 2, got %d", q.Poll())
	}
	if q.Poll() != 3 {
		t.Errorf("Expected queue poll to return 3, got %d", q.Poll())
	}
	// 测试队列的大小和空状态
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, got %d", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, got false")
	}
}

func TestDeque(t *testing.T) {
	// 创建一个空的双端队列
	d := &list.Deque{}
	// 测试双端队列的大小和空状态
	if d.Size != 0 {
		t.Errorf("Expected deque size to be 0, got %d", d.Size)
	}
	if d.Head != nil {
		t.Errorf("Expected deque head to be nil, got %v", d.Head)
	}
	if d.Tail != nil {
		t.Errorf("Expected deque tail to be nil, got %v", d.Tail)
	}
	// 从左边添加三个元素
	d.LPush(1)
	d.LPush(2)
	d.LPush(3)
	// 测试双端队列的大小和空状态
	if d.Size != 3 {
		t.Errorf("Expected deque size to be 3, got %d", d.Size)
	}
	if d.Head == nil {
		t.Errorf("Expected deque head to be not nil, got nil")
	} else if d.Head.V != 3 {
		t.Errorf("Expected deque head value to be 3, got %d", d.Head.V)
	}
	if d.Tail == nil {
		t.Errorf("Expected deque tail to be not nil, got nil")
	} else if d.Tail.V != 1 {
		t.Errorf("Expected deque tail value to be 1, got %d", d.Tail.V)
	}
	// 测试双端队列的出队操作
	if d.LPoll() != 3 {
		t.Errorf("Expected deque LPoll to return 3, got %d", d.LPoll())
	}
	if d.RPoll() != 1 {
		t.Errorf("Expected deque RPoll to return 1, got %d", d.RPoll())
	}
	if d.LPoll() != 2 {
		t.Errorf("Expected deque LPoll to return 2, got %d", d.LPoll())
	}
	// 测试双端队列的大小和空状态
	if d.Size != 0 {
		t.Errorf("Expected deque size to be 0, got %d", d.Size)
	}
	if d.Head != nil {
		t.Errorf("Expected deque head to be nil, got %v", d.Head)
	}
	if d.Tail != nil {
		t.Errorf("Expected deque tail to be nil, got %v", d.Tail)
	}
}
