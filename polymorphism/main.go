package main

import "fmt"

type PayMethod interface {
	Pay()
}

//---------------------------------------------------------------- interface

type Paypal struct{}

func (p *Paypal) Pay() {
	fmt.Println("Payment with paypal")
}

type Cash struct{}

func (c *Cash) Pay() {
	fmt.Println("Payment with cash")
}

type CreditCart struct{}

func (c *CreditCart) Pay() {
	fmt.Println("Payment with credit cart")
}

//---------------------------------------------------------------- each type implement the interface

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return &Paypal{}
	case 2:
		return &Cash{}
	case 3:
		return &CreditCart{}
	default:
		return nil
	}
}

//---------------------------------------------------------------- factory method

func main() {
	var method uint
	fmt.Println("choose payment method number")
	fmt.Println("\t 1: Paypal 2: Cash 3: CreditCart")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("error we need a valid method:")
	}
	if method > 3 {
		panic("error we need a valid method:")
	}

	p := Factory(method)
	p.Pay()
}
