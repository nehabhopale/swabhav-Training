package itemreq

import (
	"fmt"
	"machine/vendingmachine"
)

type ItemRequestedState struct {
	vendingMachine *vendingmachine.VendingMachine
}

func (i *ItemRequestedState) addItem(count uint8) error {
	fmt.Errorf("item dispense will proceed")

}
func (i *ItemRequestedState) requestItem() error {
	fmt.Errorf("item already added")

}
func (i *ItemRequestedState) insertMoney(money uint16) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	fmt.Errorf("add money")

}
