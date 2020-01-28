package main

import (
	"fmt"
)

// CreditAccount comment
type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing card payment ", amount)
}

// CreateCreditAccount comment
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}

	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

func main() {
	chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)

	chargeCh <- 500
	chargeCh <- 800

	var a string
	fmt.Scanln(&a)
}
