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

	l.size--
	lastVal := l.dataSlice[l.size]
	// 说明存在，需要删除 val 在两个数据源中的数据
	l.dataSlice[idx] = lastVal         // 将待删除元素放在末尾位置
	l.dataSlice = l.dataSlice[:l.size] // 然后删除最后一个位置
	// 如果删除的不是最后一个元素，记得更新索引
	l.dataMap[lastVal] = idx
	delete(l.dataMap, val) // 在 map 中删除
}

func (l *LRG) GetRandom() int {
	if l.size <= 0 {
		return -1
	}
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

func (i *IRG) Exist(val int) bool {
	_, ok := i.dataMap[val]
	return ok
}

type IRG struct {
	dataMap   map[int]int // <val, idx>
	dataSlice []int
	size      int
}

func NewIRG() *IRG {
	return &IRG{
		dataMap:   make(map[int]int, 10),
		dataSlice: make([]int, 0, 10),
	}
}

func (i *IRG) Insert(val int) bool {
	_, ok := i.dataMap[val]
	if ok {
		// 说明已经存在了
		return false
	}
	// 需要维护两个数据源
	i.dataMap[val] = i.size
	i.dataSlice = append(i.dataSlice, val)
	i.size++

	return true
}

func (i *IRG) Remove(val int) bool {
	idx, ok := i.dataMap[val]
	if !ok {
		// 说明没这个 val
		return false
	}

	i.size--
	lastVal := i.dataSlice[i.size]
	i.dataSlice[idx] = lastVal
	// 记得更新索引
	i.dataMap[lastVal] = idx
	// 删除对应的 val
	i.dataSlice = i.dataSlice[:i.size]
	delete(i.dataMap, val)

	return true
}

func (i *IRG) GetRandom() int {
	if i.size == 0 {
		return -1
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIdx := r.Intn(i.size) // 生成一个 [0, size) 的索引

	return i.dataSlice[randIdx]
}

var parents []int

func find(val int) int {

	for val != parents[val] {
		parent := parents[val]
		parents[val] = parents[parent]
		val = parent
	}

	return val
}
