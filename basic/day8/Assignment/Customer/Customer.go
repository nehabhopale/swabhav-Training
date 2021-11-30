package Customer

import (
	"fmt"
	Restaurant "hotel/Restaurant"
)

type Customer struct {
	name          string
	email         string
	mobile_number uint64
	password      string
	balance       uint32
	retName       string
}

func (c *Customer) GetName() string {
	return c.name
}
func (c *Customer) GetBalance() uint32 {
	return c.balance
}
func (c *Customer) SetBalance(balance uint32) {
	c.balance = balance
}

var cust []*Customer

func SignUp(newd *Customer) {
	cust = append(cust, newd)

}

func (c *Customer) PlaceOrder(rest *Restaurant.Restaurant, dish string) {
	var bill uint32 = 0
	orderPrize, availability := rest.GetItemDetails(dish)
	if availability != false {
		fmt.Println("I would like to order", dish)
		bill += orderPrize
		if int(bill) > int(c.balance) {
			fmt.Println("sorry u don't have suffiecient balance to place an order")
		} else {
			fmt.Println("your order placed successfully")
			b := c.GetBalance() - bill
			c.SetBalance(b)
		}
	} else {
		fmt.Println(dish, " is currently unavailable ")
	}

}
func New(name string, email string, mobileNum uint64, pwd string, balance uint32) *Customer {
	return &Customer{
		name:          name,
		email:         email,
		mobile_number: mobileNum,
		password:      pwd,
		balance:       balance,
	}
}