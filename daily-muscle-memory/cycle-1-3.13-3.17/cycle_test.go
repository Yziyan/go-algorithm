// @Author: Ciusyan 3/13/24

package cycle_1_3_13_3_17

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func TestLRG(t *testing.T) {
	lrg := NewLRG()

	// 1.测试 Insert
	lrg.Insert(2)
	lrg.Insert(3)
	lrg.Insert(1)
	lrg.Insert(2)
	lrg.Insert(5)

	assert.True(t, lrg.Exist(1))
	assert.True(t, lrg.Exist(2))
	assert.True(t, lrg.Exist(3))
	assert.True(t, !lrg.Exist(4))
	assert.True(t, lrg.Exist(5))
	assert.True(t, !lrg.Exist(6))

	// 2.测试 Random
	for i := 0; i < 4; i++ {
		// 这个只能主观臆断了
		t.Log(lrg.GetRandom())
	}

	// 3.测试 Remove
	lrg.Remove(2)
	lrg.Remove(2)
	lrg.Remove(3)
	lrg.Remove(4)
	lrg.Remove(5)
	lrg.Remove(6)

	assert.True(t, !lrg.Exist(2))
	assert.True(t, !lrg.Exist(3))
	assert.True(t, !lrg.Exist(4))
	assert.True(t, !lrg.Exist(5))
	assert.True(t, !lrg.Exist(6))

	lrg.Remove(1)
	assert.True(t, !lrg.Exist(1))
	lrg.Insert(1)
	assert.True(t, lrg.Exist(1))

	t.Run("case", func(t *testing.T) {
		set := Constructor()
		assert.True(t, !set.Remove(0))
		assert.True(t, !set.Remove(0))
		assert.True(t, set.Insert(0))
		assert.True(t, set.GetRandom() == 0)
		assert.True(t, set.Remove(0))
		assert.True(t, set.Insert(0))
	})

	t.Run("case2", func(t *testing.T) {
		set := Constructor()
		t.Log(set.GetRandom())
		num := 10000
		for i := 0; i < num; i++ {
			assert.True(t, set.Insert(i))
		}

		for i := 0; i < num; i += 2 {
			assert.True(t, set.Remove(i))
		}

	})

}

func TestIRG(t *testing.T) {

	irg := NewIRG()

	assert.True(t, irg.Insert(1))
	assert.True(t, !irg.Insert(1))
	assert.True(t, irg.Insert(2))
	assert.True(t, !irg.Insert(2))
	assert.True(t, !irg.Insert(2))
	assert.True(t, irg.Insert(3))
	assert.True(t, irg.Insert(4))
	assert.True(t, irg.Insert(5))
	assert.True(t, irg.Insert(6))
	for i := 0; i < 3; i++ {
		t.Log(irg.GetRandom())
	}

	assert.True(t, irg.Exist(1))
	assert.True(t, irg.Exist(2))
	assert.True(t, irg.Exist(3))
	assert.True(t, irg.Exist(4))
	assert.True(t, irg.Exist(5))
	assert.True(t, irg.Exist(6))
	assert.True(t, !irg.Exist(8))
	assert.True(t, !irg.Exist(10))

	assert.True(t, irg.Remove(1))
	assert.True(t, !irg.Exist(1))

	assert.True(t, !irg.Remove(8))
	assert.True(t, !irg.Remove(1))

	assert.True(t, irg.Remove(2))
	assert.True(t, !irg.Exist(2))
	assert.True(t, !irg.Remove(2))
	assert.True(t, irg.Insert(2))
	assert.True(t, irg.Exist(2))
	assert.True(t, irg.Exist(4))
	assert.True(t, irg.Exist(5))

	for i := 0; i < 3; i++ {
		t.Log(irg.GetRandom())
	}
}

func findMaxNumber(n int, nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	maxLen := len(strconv.Itoa(n))
	var dfs func(current int, depth int) int
	dfs = func(current int, depth int) int {
		if depth > maxLen {
			return -1 // 超过最大长度，剪枝
		}
		if current >= n {
			return -1 // 超过n，剪枝
		}
		maxResult := current
		for _, num := range nums {
			next := current*10 + num
			result := dfs(next, depth+1)
			if result > maxResult {
				maxResult = result
			}
		}
		return maxResult
	}

	result := dfs(0, 0)
	if result == 0 {
		return -1 // 如果没有找到合适的数，返回-1
	}
	return result
}

func TestFindMaxNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		nums []int
		want int
	}{
		{
			name: "Test 1",
			n:    100,
			nums: []int{1, 2, 3},
			want: 33,
		},
		{
			name: "Test 2",
			n:    200,
			nums: []int{1, 3, 5},
			want: 155,
		},
		{
			name: "Test 3",
			n:    50,
			nums: []int{5, 7, 8},
			want: 8, // 在此案例中，没有比n小的数可以由给定数组组成
		},
		{
			name: "Test 4",
			n:    1000,
			nums: []int{0, 1, 2, 3},
			want: 333,
		},
		{
			name: "Test 5",
			n:    432,
			nums: []int{4, 2, 1},
			want: 424,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxNumber(tt.n, tt.nums); got != tt.want {
				t.Errorf("findMaxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
