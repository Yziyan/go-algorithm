// @Author: Ciusyan 3/14/24

package cycle_1_3_13_3_17

import (
	"math/rand"
	"time"
)

/**
时间复杂度 O(1) 的限制下实现 Insert(val)，Remove(val)，GetRandom() 三个函数，
GetRandom() 要求每一个值被返回的概率相等。
*/

type LRG struct {
	// 存储值和对应切片中的索引，<3, 1> 代表 val = 3，idx = 1
	dataMap   map[int]int
	dataSlice []int
	size      int
}

func NewLRG() *LRG {
	return &LRG{
		dataMap:   make(map[int]int, 10),
		dataSlice: make([]int, 0, 10),
	}
}

func (l *LRG) Insert(val int) {
	_, ok := l.dataMap[val]
	if ok {
		// 说明这个值以前存在，直接返回
		return
	}

	// 否则说明第一次插入，写入两个依赖，并更新 size
	l.dataMap[val] = l.size
	l.dataSlice = append(l.dataSlice, val)
	l.size++
}

func (l *LRG) Remove(val int) {
	idx, ok := l.dataMap[val]
	if !ok {
		// 说明都不存在
		return
	}

	// 说明存在，需要删除 val 在两个数据源中的数据
	delete(l.dataMap, val)                   // 在 map 中删除
	l.dataSlice[idx] = l.dataSlice[l.size-1] // 将待删除元素放在末尾位置
	l.dataSlice = l.dataSlice[:l.size-1]     // 然后删除最后一个位置
	l.size--
	if l.size != 0 {
		// 如果删除的不是最后一个元素，记得更新索引
		l.dataMap[l.dataSlice[idx]] = idx
	}
}

func (l *LRG) GetRandom() int {
	// 防止每次生成的随机数相同，设置随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 先随机从切片中获取一个出来
	idx := r.Intn(l.size) // 生成一个 [0, size) 的数

	// 返回对应索引位置的元素
	return l.dataSlice[idx]
}

func (l *LRG) Exist(val int) bool {
	_, ok := l.dataMap[val]
	return ok
}
