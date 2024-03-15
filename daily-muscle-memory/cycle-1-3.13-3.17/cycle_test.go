// @Author: Ciusyan 3/13/24

package cycle_1_3_13_3_17

import (
	"github.com/stretchr/testify/assert"
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
