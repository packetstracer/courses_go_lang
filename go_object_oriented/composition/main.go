package main

import (
	"fmt"
)

// Account comment
type Account struct{}

// ProcessPayment comment
func (a *Account) ProcessPayment(amount float32) {
	fmt.Println("Processing Payment... ", amount)
}

// AvailableFunds comment
func (a *Account) AvailableFunds() int {
	fmt.Println("Returning avaiable funds")

	return 0
}

// CreditAccount comment
type CreditAccount struct {
	Account
}

func main() {
	c := &CreditAccount{}

	c.ProcessPayment(100)
	c.ProcessPayment(300)
	c.AvailableFunds()
}
