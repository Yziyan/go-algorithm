// @Author: Ciusyan 10/30/23

package other

import (
	"fmt"
	"sync/atomic"
	"time"
)

type MiniSema struct {
	// 最大允许的并发数量
	maxConcurrent int32
	// 正在运行的协程数量
	cur int32
}

type MiniSemaOptions func(sema *MiniSema)

// WithMaxConcurrentOption 可以选择传入支持的最大并发数
func WithMaxConcurrentOption(max int32) MiniSemaOptions {
	return func(sema *MiniSema) {
		sema.maxConcurrent = max
	}
}

// NewMiniSema 初始请传入最大并发的数量
func NewMiniSema(opts ...MiniSemaOptions) *MiniSema {
	// 默认最大的并发量是 1
	sema := &MiniSema{
		maxConcurrent: 1,
	}

	for _, opt := range opts {
		opt(sema)
	}

	return sema
}

// UpdateMaxConcurrent 修改能最大运行的并发数量
func (s *MiniSema) UpdateMaxConcurrent(newMaxConcurrent int32) {
	// 需要使用原子操作
	atomic.StoreInt32(&s.maxConcurrent, newMaxConcurrent)
}

// GetMaxConcurrent 获取最大并发执行数
func (s *MiniSema) GetMaxConcurrent() int32 {
	return atomic.LoadInt32(&s.maxConcurrent)
}

// TryAcquire 尝试获取一个信号量，能获取到就返回 ture，不能就返回 false
func (s *MiniSema) TryAcquire() bool {
	cur := atomic.LoadInt32(&s.cur)
	mx := atomic.LoadInt32(&s.maxConcurrent)
	if cur >= mx {
		return false
	}

	// 说明可以获取
	cur = atomic.AddInt32(&s.cur, 1)
	// 这里需要做 double check
	if cur > mx {
		// 说明有别的协程刚刚增加了，把刚刚增加的减回来
		atomic.AddInt32(&s.cur, -1)
		return false
	}

	return true
}

// Release 释放 n 个信号量
func (s *MiniSema) Release(n int32) {
	cur := atomic.AddInt32(&s.cur, -n)
	if cur < 0 {
		// 如果减完都小于 0 了，说明外界传参有问题，本来是用户的问题，那么用户就得承受并发数量的控制有问题。
		// 但是我们稍微守一个底线就是不能小于 0
		atomic.StoreInt32(&s.cur, 0)
	}
}

// ListenMaxConcurrent 在后台监听
//	getValFunc 告诉如何获取值，duration 多久轮询一次
//		// 默认使用 1 个并发数
//		miniSema = sema.NewMiniSema()
//		// 开启监听，一分钟监听一次
//		miniSema.ListenMaxConcurrent(func() int32 {
//			// 从 apollo 获取最大的并发数量
//			conf := apollox.GetBatchPushMaterialConf()
//			if conf == nil {
//				return 0
//			}
//			return conf.MaxConcurrent
//		}, time.Minute)
func (s *MiniSema) ListenMaxConcurrent(getValFunc func() int32, duration time.Duration) {
	go func() {
		// 协程内部必须捕获异常
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		// 一直轮询监听，间隔为 duration
		for {
			newVal := getValFunc()
			oldMax := s.GetMaxConcurrent()
			if newVal > 0 && newVal != oldMax {
				s.UpdateMaxConcurrent(newVal)
				fmt.Printf("修改成功：old = %d, new = %d\n", oldMax, newVal)
			}
			time.Sleep(duration)
		}
	}()
}
