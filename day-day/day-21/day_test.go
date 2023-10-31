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

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
			b[i] = charset[r.Intn(len(charset))]
		}
		return string(b)
	}

	randStringSlice := func(size, strLen int) []string {
		res := make([]string, size)
		for i := 0; i < size; i++ {
			res[i] = randString(strLen)
		}
		return res
	}

	loopCount := 1000
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
		size = r.Intn(size) + 1
		res := make([][2]int, size)

		for i := 0; i < size; i++ {
			var (
				beginTime = r.Intn(maxTime) + 1
				endTime   = r.Intn(maxTime) + 1
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

	heap := NewHeap[int](func(x, y int) int {
		return x - y
	})

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

func TestLessMoneySplitGold(t *testing.T) {

	randArr := func(size int, max int) []int {
		size = r.Intn(size) + 1
		res := make([]int, size)

		for i := 0; i < size; i++ {
			res[i] = r.Intn(max) + 1
		}

		return res
	}

	arrLen := 20
	mx := 100
	loopSize := 10000

	for i := 0; i < loopSize; i++ {
		arr := randArr(arrLen, mx)
		arr1 := slices.Clone(arr)

		got := LessMoneySplitGold(arr)
		got1 := LessMoneySplitGold(arr1)

		assert.Equal(t, got, got1)
	}

}

func TestFindMaximizedCapital(t *testing.T) {

	profit := findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2})

	t.Log(profit)
}

func TestMinLight(t *testing.T) {

	testCases := []struct {
		name string
		road string
		want int
	}{
		{"空路", "", 0},
		{"只有墙", "X", 0},
		{"只有一个街道", ".", 1},
		{"一个街道两边是墙", "X.X", 1},
		{"连续三个街道", "...", 1},
		{"两个街道两边是墙", "X..X", 1},
		{"三个街道两边是墙", "X...X", 1},
		{"中间有一个墙", "..X..", 2},
		{"中间有一个墙，两边各有三个街道", "...X...", 2},
		{"中间有一个墙，两边各有四个街道", "....X....", 4},
		{"多个分段的街道", "XX...XX...XX", 2},
		{"复杂街道", "..X...X....X.....X..X...X..X...X..", 11},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MinLight(tc.road)
			assert.Equal(t, tc.want, got)
		})
	}

}
