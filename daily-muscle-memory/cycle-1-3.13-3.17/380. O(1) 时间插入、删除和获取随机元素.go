// @Author: Ciusyan 3/14/24

package cycle_1_3_13_3_17

import (
	"math/rand"
	"time"
)

/* 思路重复，伪代码
// 分析
1.既然要使用 O(1) 操作实现三个操作，那么要能联想到：数组的随机访问、哈希表的 CRUD
2.再看看要实现的三个函数，增删、随机查，那么数据源需要使用：map+slice
2.1 增：map 增加 O(1)，append 至 slice 末尾 O(1) ps: 不扩容的情况下
2.2 删：map 删除 O(1)，先找出删除元素的索引，然后交换至末尾，再删除末尾的元素
2.3 随机查：生成一个 [0, size) 的索引，直接在 slice 中返回。

所以对于增删改查的操作，使用 map+slice 都可以完成，但是对于删除，
我们需要找到对应 val 的索引，所以 map 的 val 就存储对应 slice 的索引

那么在做相应变动的时候，可能需要做两个数据源的变更

// 思路/伪代码
struct： dataMap、dataSlice、size(可选)
Insert：不存在 val 才添加。并且添加时要维护两个数据源，插入 map 和 追加至 slice 末尾
Remove: 存在 val 才删除，先利用 map 找出 val 所在位置的 idx，
		然后将末尾元素交换过来，删除 slice 末尾的元素，记得更新末尾元素在 map 中的索引
Random：直接生成一个 [0, size) 范围内的 idx，然后直接从 slice 中取出返回。
*/

type RandomizedSet struct {
	dataMap   map[int]int // <val, slice.val.idx>
	dataSlice []int
	size      int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		dataMap:   make(map[int]int, 10),
		dataSlice: make([]int, 0, 10),
		size:      0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dataMap[val]
	if ok {
		// 说明已经存在了
		return false
	}

	// 不存在才插入，两个数据源
	this.dataMap[val] = this.size
	this.dataSlice = append(this.dataSlice, val)
	this.size++

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dataMap[val]
	if !ok {
		return false
	}

	// 肯定需要删除了
	this.size--
	// 删除 slice 中的数据
	lastVal := this.dataSlice[this.size]
	this.dataSlice[idx] = lastVal
	// 但是别忘记更新最后一个元素在 slice 中的索引
	this.dataMap[lastVal] = idx
	this.dataSlice = this.dataSlice[0:this.size]

	// 删除 map 中的数据
	delete(this.dataMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if this.size == 0 {
		return -1
	}
	// 设置随机种子，防止每次生成的随机数相同
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成 [0, size) 的索引
	randIdx := r.Intn(this.size)

	return this.dataSlice[randIdx]
}
