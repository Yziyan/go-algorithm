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
}
