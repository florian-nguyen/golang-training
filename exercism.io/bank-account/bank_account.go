package account

/*
Simulate a bank account supporting opening/closing, withdrawals, and deposits
of money. Watch out for concurrent transactions!

A bank account can be accessed in multiple ways. Clients can make
deposits and withdrawals using the internet, mobile phones, etc. Shops
can charge against the account.

Create an account that can be accessed from multiple threads/processes
(terminology depends on your programming language).

It should be possible to close an account; operations against a closed
account must fail.
*/

import (
	"sync"
)

// Account definition
type Account struct {
	sync.Mutex // put before monitored parameters
	balance    int64
	closed     bool
}

// Open creates an account with a given starting amount
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

// Close will empty and close the account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	// exit if already closed
	if a.closed {
		return 0, false
	}

	payout = a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance gives the current amount of money in the account
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit modifies the amount of money on the account and returns the new balance value
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	// Account already closed or invalid deposit should return (0, false)
	if a.closed || amount+a.balance < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}
