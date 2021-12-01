package hasitem

import (
	"fmt"
	"machine/vendingmachine"
)

type HasItemState struct {
	vendingMachine *vendingmachine.VendingMachine
}

//we r having item and we r requesting it
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requestd by customer")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}
func (i *HasItemState) addItem(count uint8) error {
	fmt.Println(count, "no of items added")
	i.vendingMachine.incrementItemCount(count)
	return nil
}
func (i *HasItemState) insertMoney(money uint16) error {
	fmt.Errorf("u will need to select item")
}
func (i *HasItemState) dispenseItem() error {
	fmt.Errorf("u will need to select item to be dispended")
}
