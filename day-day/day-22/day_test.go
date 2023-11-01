// @Author: Ciusyan 10/31/23

package day_22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()

	// 建集合
	uf.MakeSets(1, 2, 3, 4, 5)

	// 测试 IsSame
	t.Run("IsSame", func(t *testing.T) {
		assert.True(t, uf.IsSame(1, 1))
		assert.False(t, uf.IsSame(1, 2))
	})

	// 测试 Union
	t.Run("Union", func(t *testing.T) {
		// 将 1 和 2 合并到同一个集合
		uf.Union(1, 2)
		assert.True(t, uf.IsSame(1, 2))

		// 将 3 和 4 合并到同一个集合
		uf.Union(3, 4)
		assert.True(t, uf.IsSame(3, 4))

		// 再将这两个集合合并
		uf.Union(1, 3)
		assert.True(t, uf.IsSame(1, 4))
	})

	// 测试 Find
	t.Run("Find", func(t *testing.T) {
		root := uf.Find(1)
		assert.Equal(t, root, 4)
	})

	// 继续建集合
	uf.MakeSets(6, 7, 8, 9, 10)
	t.Run("Rank", func(t *testing.T) {
		uf.Union(6, 7)
		assert.True(t, uf.IsSame(6, 7))

		uf.Union(7, 8)
		assert.True(t, uf.IsSame(6, 8))
		assert.Equal(t, uf.Find(6), 7)
		assert.Equal(t, uf.Find(7), 7)
		assert.Equal(t, uf.Find(8), 7)
	})

	t.Run("依赖上面的执行", func(t *testing.T) {
		uf.Union(1, 8)
		assert.True(t, uf.IsSame(4, 7))
		assert.True(t, uf.IsSame(3, 8))
		assert.Equal(t, uf.Find(7), 4)
		assert.Equal(t, uf.Find(8), 4)
		assert.Equal(t, uf.Find(1), 4)
		assert.Equal(t, uf.Find(3), 4)
	})

	t.Log("over")
}

func TestFindCircleNum(t *testing.T) {
	args := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	num := findCircleNum(args)
	t.Log(num)
}
