// @Author: Ciusyan 10/16/23

package other

import (
	"sync"
	"testing"
)

// 测试锁和同步原语
// 需求：给一个账户转账，每次转 1 元，但是要分 10000 次转

func TestLockTransfer(t *testing.T) {

}

type Account struct {
	accountId int
	balance   int

	mu sync.Mutex
}

func NewAccount(accountId int) *Account {
	return &Account{
		accountId: accountId,
	}
}

func (a *Account) Add(amount int) {
	a.balance += amount
}

func (a *Account) AddLock(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}

func (a *Account) AddCAP() {

}

// 给账户 account 转账 amount ￥
func TransferLock(account *Account, amount int) {
	account.Add(amount)
}
