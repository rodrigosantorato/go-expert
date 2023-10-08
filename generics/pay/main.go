package main

import "fmt"

func main() {
	c := &cash{}
	card := &creditCard{}
	processPurchase(card)
	processPurchase(c)
}

type creditCard struct {
}

func (c *creditCard) Pay() bool {
	return true
}

func (c *creditCard) Name() string {
	return "credit card"
}

type cash struct {
}

func (c *cash) Pay() bool {
	return true
}

func (c *cash) Name() string {
	return "cash"
}

type paymentMethod interface {
	Pay() bool
	Name() string
}

func processPurchase[T paymentMethod](m T) bool {
	fmt.Printf("paying with %v\n", m.Name())
	return m.Pay()
}
