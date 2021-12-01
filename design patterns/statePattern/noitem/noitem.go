package noitem

import (
	"fmt"
	"machine/vendingmachine"
)

type NoItem struct {
	vendingMachine *vendingmachine.VendingMachine
}

func (i *NoItem) requestItem() error {
	return fmt.Errorf("Item is not available")
}

func (i *NoItem) addItem(count uint8) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItem) insertMoney(money uint16) error {
	return fmt.Errorf("Sorry Item is not available")
}
func (i *NoItem) dispenseItem() error {
	return fmt.Errorf("sorry item is not available so cant be dispensed")
}
