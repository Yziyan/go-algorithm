// @Author: Ciusyan 9/22/23

package day_15

import (
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()

	// 添加元素
	heap.Add(5)
	heap.Add(7)
	heap.Add(3)
	heap.Add(10)
	heap.Add(2)

	// 获取堆顶元素
	top := heap.Get()
	if top != 2 {
		t.Errorf("Expected top element to be 2, but got %d", top)
	}

	// 删除堆顶元素
	removed := heap.Remove()
	if removed != 2 {
		t.Errorf("Expected removed element to be 2, but got %d", removed)
	}

	// 替换堆顶元素
	replaced := heap.Replace(8)
	if replaced != 3 {
		t.Errorf("Expected replaced element to be 3, but got %d", replaced)
	}

	// 再次获取堆顶元素
	newTop := heap.Get()
	if newTop != 5 {
		t.Errorf("Expected new top element to be 5, but got %d", newTop)
	}

	// 清空堆
	heap.Clear()
	if heap.Size() != 0 || !heap.IsEmpty() {
		t.Errorf("Expected heap to be empty after clearing, but got size %d", heap.Size())
	}
}

func TestSortedLengthK(t *testing.T) {
	nums := []int{3, 4, 1, 5, 4, 6, 8, 6}
	sortedArrLengthK(nums, 2)
	t.Log(nums)
}
