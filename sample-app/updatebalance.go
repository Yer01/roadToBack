package main

import (
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
	if t.transactionType != transactionWithdrawal && t.transactionType != transactionDeposit {
		return fmt.Errorf("unknown transaction type")
	}
	if t.transactionType == transactionDeposit {
		c.balance += t.amount
	} else {
		if t.amount > c.balance {
			return fmt.Errorf("insufficient funds")
		}
		c.balance -= t.amount
	}
	return nil
}
