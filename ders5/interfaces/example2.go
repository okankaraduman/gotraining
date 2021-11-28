package main

import "fmt"

func main() {
	x := new(Account)
	x.AvailableFunds()
	x.ProcessPayment(20)

	y := new(CreditAccount)
	y.ID = 0
}

type Account struct {
	ID int
}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("listing Available funds")
	return 0
}
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing Payment")
	return true
}

type CreditAccount struct {
	Account
}
