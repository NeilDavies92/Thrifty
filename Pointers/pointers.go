package main

import (
	"fmt"
)

// Go supports pointers, allowing you to pass references to values
// and records within your program.

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value.
type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {

}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
