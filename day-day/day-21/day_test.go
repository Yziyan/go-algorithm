// @Author: Ciusyan 10/25/23

package day_21

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	idx := 3
	strs := []int{44, 22, 133, 765, 43}
	strs = append(strs[0:idx], strs[idx+1:]...)
	fmt.Println(strs)
}

func TestLowestString(t *testing.T) {

	randString := func(strLen int) string {
		const charset = "abcdefghijklmnopqrstuvwxyz"
		b := make([]byte, strLen)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}

	randStringSlice := func(size, strLen int) []string {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		res := make([]string, size)
		for i := 0; i < size; i++ {
			res[i] = randString(strLen)
		}
		return res
	}

	loopCount := 50
	arrLen := 6
	strLen := 4
	for i := 0; i < loopCount; i++ {
		strs := randStringSlice(arrLen, strLen)
		t.Log(strs)
		str2 := slices.Clone(strs)
		got1 := LowestString(strs)
		got2 := LowestString2(str2)
		assert.Equal(t, got1, got2)
		t.Log(got1)
		t.Log(got2)
		fmt.Println()
	}

	strs := []string{"b", "ba"}
	got1 := LowestString(strs)

	strs = []string{"b", "ba"}
	got2 := LowestString2(strs)

	t.Log(got1)
	t.Log(got2)
	assert.Equal(t, got1, got2)

}

func TestBestArrange(t *testing.T) {

	// 随机生成会议
	randArranges := func(size int, maxTime int) [][2]int {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		size = rand.Intn(size) + 1
		res := make([][2]int, size)

		for i := 0; i < size; i++ {
			var (
				beginTime = rand.Intn(maxTime) + 1
				endTime   = rand.Intn(maxTime) + 1
			)

			if beginTime == endTime {
				endTime++
			} else if beginTime > endTime {
				// 说明开始比结束还晚，交换一下
				beginTime, endTime = endTime, beginTime
			}

			res[i] = [2]int{beginTime, endTime}
		}

		return res
	}

	arrangesLen := 20
	maxTime := 30
	loopSize := 10000

	for i := 0; i < loopSize; i++ {
		arranges := randArranges(arrangesLen, maxTime)
		arranges1 := slices.Clone(arranges)

		got := BestArrange(arranges)
		got1 := BestArrange1(arranges1)

		assert.Equal(t, got, got1)
	}
}

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

	// 再次获取堆顶元素
	newTop := heap.Get()
	if newTop != 3 {
		t.Errorf("Expected new top element to be 3, but got %d", newTop)
	}

	// 删除堆顶元素
	removed = heap.Remove()
	if removed != 3 {
		t.Errorf("Expected removed element to be 3, but got %d", removed)
	}

	newTop = heap.Get()
	if newTop != 5 {
		t.Errorf("Expected new top element to be 5, but got %d", newTop)
	}
}
