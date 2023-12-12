// @Author: Ciusyan 12/12/23

package day_27

// DoubleQueue 用切片模拟一个双端队列
type DoubleQueue []int

func NewDoubleQueue() DoubleQueue {
	return make(DoubleQueue, 0)
}

func (q *DoubleQueue) Size() int {
	return len(*q)
}

// OfferLeft 从左侧入队
func (q *DoubleQueue) OfferLeft(val int) {
	temp := []int{val}
	*q = append(temp, *q...)
}

// OfferRight 从右侧入队
func (q *DoubleQueue) OfferRight(val int) {
	*q = append(*q, val)
}

// PollLeft 从左侧出队
func (q *DoubleQueue) PollLeft() int {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

// PollRight 从右侧出队
func (q *DoubleQueue) PollRight() int {
	last := len(*q) - 1
	res := (*q)[last]
	*q = (*q)[:last]
	return res
}
