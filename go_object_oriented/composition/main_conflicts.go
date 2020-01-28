package main

import (
	"fmt"
)

// CreditAccount comment
type CreditAccount struct{}

// ProcessPayment comment
func (ca *CreditAccount) ProcessPayment(amount float32) {
	fmt.Println("Processing Payment... ", amount)
}

// AvailableFunds comment
func (ca *CreditAccount) AvailableFunds() int {
	fmt.Println("Returning Credit Account available funds")

	return 100
}

// CheckingAccount comment
type CheckingAccount struct{}

// AvailableFunds comment
func (ch *CheckingAccount) AvailableFunds() int {
	fmt.Println("Returning Checking Account avaiable funds")

	return 200
}

// Account comment
type Account struct {
	CreditAccount
	CheckingAccount
}

// AvailableFunds comment
func (a *Account) AvailableFunds() int {
	fmt.Println("Returning Account avaiable funds")

	return a.CreditAccount.AvailableFunds() + a.CheckingAccount.AvailableFunds()
}

func main() {
	c := &Account{}

	c.ProcessPayment(100)
	c.ProcessPayment(300)

	fmt.Println(c.AvailableFunds())
}
