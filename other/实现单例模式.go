// @Author: Ciusyan 2023/7/29

package other

import "sync"

// 不用导出
type singleton struct {
}

// 饿汉式，直接创建好一个实例
var instance1 = &singleton{}

func GetInstance1() *singleton {
	return instance1
}

// 懒汉式，第一次使用的时候才创建，但是可能会并发不安全
var instance *singleton

func GetInstance() *singleton {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		// 初始化实例
		instance = &singleton{}
	}

	return instance
}

func GetInstance2() *singleton {
	var oc sync.Once
	oc.Do(func() {
		// 初始化实例
		instance = &singleton{}
	})

	return instance
}
