// @Author: Ciusyan 6/19/24

package phase3

import (
	"math/rand"
	"time"
)

// https://leetcode.cn/problems/insert-delete-getrandom-o1/

type RandomizedSet struct {
	mapEle   map[int]int
	sliceEle []int
	size     int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapEle:   make(map[int]int, 10),
		sliceEle: make([]int, 0, 10),
		size:     0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.mapEle[val]
	if ok {
		return false
	}

	this.sliceEle = append(this.sliceEle, val)
	this.mapEle[val] = this.size

	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.mapEle[val]
	if !ok {
		return false
	}
	this.size--
	last := this.sliceEle[this.size]
	this.sliceEle[idx] = last
	this.mapEle[last] = idx
	this.sliceEle = this.sliceEle[:this.size]
	delete(this.mapEle, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rIdx := r.Intn(this.size)
	return this.sliceEle[rIdx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
