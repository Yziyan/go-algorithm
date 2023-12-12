// @Author: Ciusyan 12/12/23

package day_27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleQueue(t *testing.T) {
	queue := NewDoubleQueue()
	assert.True(t, queue.Size() == 0)
	queue.OfferRight(3)
	queue.OfferRight(4)
	queue.OfferRight(5)
	assert.True(t, queue.Size() == 3)
	assert.True(t, queue.PollLeft() == 3)
	assert.True(t, queue.PollRight() == 5)
	assert.True(t, queue.Size() == 1)
	queue.OfferLeft(10)
	queue.OfferLeft(20)
	assert.True(t, queue.PollLeft() == 20)
	assert.True(t, queue.Size() == 2)
}
