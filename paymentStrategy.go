package main

import "fmt"


	type PaymentStrategy interface {
		Pay(amount float64)
	}

	type CashPayment struct{}

	func (c *CashPayment) Pay(amount float64) {
		fmt.Printf("Paid %.2f using Credit Card.\n", amount)
	}

	type KaspiPayment struct{}

	func (k *KaspiPayment) Pay(amount float64) {
		fmt.Printf("Paid %.2f using Kaspi.\n", amount)
	}
	

