// @Author: Ciusyan 10/16/23

package other

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 测试锁和同步原语
// 需求：给一个账户转账，每次转 amount 元，但是要分 loopCount 次转

func TestLockTransfer(t *testing.T) {

	loopCount := int64(10000000)
	amount := int64(1)
	testCases := []struct {
		name string

		account   *Account
		loopCount int64  // 循环的次数
		amount    int64  // 每次加多少钱
		md        method // 使用的方法

		wantBalance int64
		wantFunc    func(want, got int64) (string, bool)
	}{
		{
			name:    "【常规】",
			account: NewAccount(1),
			md:      regular,

			wantBalance: amount * loopCount,
			wantFunc: func(want, got int64) (string, bool) {
				// 这里大概率是不会到达 wantBalance 的，我们默认是不可能事件了
				if want == got {
					return fmt.Sprintf("得到了 got: %d, 不可能到达 want: %d", got, want), false
				}

				return "", true
			},
		},
		{
			name:    "【带锁】",
			account: NewAccount(2),
			md:      lock,

			wantBalance: amount * loopCount,
			wantFunc: func(want, got int64) (string, bool) {
				// 这里大概率是不会到达 wantBalance 的，我们默认是不可能事件了
				if want != got {
					return fmt.Sprintf("实际得到了 got: %d, 想要到达 want: %d", got, want), false
				}

				return "", true
			},
		},
		{
			name:    "【原子操作-CAS】",
			account: NewAccount(3),
			md:      cas,

			wantBalance: amount * loopCount,
			wantFunc: func(want, got int64) (string, bool) {
				// 这里大概率是不会到达 wantBalance 的，我们默认是不可能事件了
				if want != got {
					return fmt.Sprintf("实际得到了 got: %d, 想要到达 want: %d", got, want), false
				}

				return "", true
			},
		},
		{
			name:    "【原子操作-FAA】",
			account: NewAccount(4),
			md:      faa,

			wantBalance: amount * loopCount,
			wantFunc: func(want, got int64) (string, bool) {
				// 这里大概率是不会到达 wantBalance 的，我们默认是不可能事件了
				if want != got {
					return fmt.Sprintf("实际得到了 got: %d, 想要到达 want: %d", got, want), false
				}

				return "", true
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			wg := sync.WaitGroup{}
			start := time.Now()
			for i := 0; i < int(loopCount); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					TransferLock(tt.account, amount, tt.md)
				}()
			}

			wg.Wait()
			elapsed := time.Since(start)
			t.Log(tt.account.info())
			t.Logf("%s执行耗时: %s\n", tt.name, elapsed)

			msg, ok := tt.wantFunc(tt.wantBalance, tt.account.balance)
			assert.True(t, ok, msg)
		})
	}
}

type Account struct {
	accountId int
	balance   int64

	mu sync.Mutex
}

func NewAccount(accountId int) *Account {
	return &Account{
		accountId: accountId,
	}
}

func (a *Account) info() string {
	return fmt.Sprintf("账户：%d，余额为：%d", a.accountId, a.balance)
}

func (a *Account) Add(amount int64) {
	a.balance += amount
}

func (a *Account) AddLock(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}

func (a *Account) AddCAS(amount int64) {
	for {
		old := a.balance
		if atomic.CompareAndSwapInt64(&a.balance, old, old+amount) {
			// 如果操作成功了，就退出
			break
		}
	}
}

func (a *Account) AddFAA(amount int64) {
	// 直接使用 FAA
	atomic.AddInt64(&a.balance, amount)
}

type method string

const (
	regular = method("regular")
	lock    = method("lock")
	cas     = method("cas")
	faa     = method("faa")
)

// 给账户 account 转账 amount ￥
func TransferLock(account *Account, amount int64, md method) {
	switch md {
	case regular:
		account.Add(amount)
	case lock:
		account.AddLock(amount)
	case cas:
		account.AddCAS(amount)
	case faa:
		account.AddFAA(amount)
	default:
		panic("不支持此方法")
	}
}

var counter int32 = 0

var max int32 = 1 // 最大允许的并发数量

// 在运行时更新 max 的值
func updateMax(max *int32, newMaxConcurrent int32) {
	atomic.StoreInt32(max, newMaxConcurrent)
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup

	// 模拟修改成 3
	go func() {
		time.Sleep(2 * time.Second)
		t.Log("max修改为3")
		updateMax(&max, 3)
	}()

	// 模拟修改成 1
	go func() {
		time.Sleep(6 * time.Second)
		t.Log("max修改为1")
		updateMax(&max, 1)
	}()

	// 模拟改成 5
	go func() {
		time.Sleep(10 * time.Second)
		t.Log("max修改为5")
		updateMax(&max, 5)
	}()

	// 模拟改成 5
	go func() {
		time.Sleep(12 * time.Second)
		t.Log("max修改为2")
		updateMax(&max, 2)
	}()

	for i := 1; i <= 150; i++ {
		wg.Add(1)
		go func(taskId int) {
			defer wg.Done()

			if !acquireSemaphore() {
				// 没获取到锁
				t.Logf("TaskId = %d 尚未获取到锁", taskId)
				return
			}
			// 获取信号量
			defer releaseSemaphore() // 释放信号量

			// 执行任务
			t.Logf("TaskID = %d 开始执行\n", taskId)
			time.Sleep(2 * time.Second)
			t.Logf("TaskID = %d 执行结束\n", taskId)
		}(i)

		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

func acquireSemaphore() bool {
	n := atomic.LoadInt32(&counter)
	mx := atomic.LoadInt32(&max)
	if n >= mx {
		return false
	}

	// 说明可以获取
	n = atomic.AddInt32(&counter, 1)
	if n > mx {
		// 说明有别的协程刚刚增加了，把刚刚增加的减回来
		n = atomic.AddInt32(&counter, -1)

		return false
	}

	return true
}

func releaseSemaphore() {
	atomic.AddInt32(&counter, -1)
}
