package main

import (
	"fmt"
	"log"
)

type State interface {
	addItem(uint8) error
	requestItem() error
	insertMoney(money uint16) error
	dispenseItem() error
}
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount uint8
	itemPrice uint16
}

func New(itemCount uint8, itemPrice uint16) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItem{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count uint8) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money uint16) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count uint8) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

///////
type NoItem struct {
	vendingMachine *VendingMachine
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

///
type HasItemState struct {
	vendingMachine *VendingMachine
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
	return fmt.Errorf("u will need to select item")
}
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("u will need to select item to be dispended")
}

////
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count uint8) error {
	return fmt.Errorf("item dispense will proceed")

}
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already added")

}
func (i *ItemRequestedState) insertMoney(money uint16) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("add money")

}

//
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count uint8) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money uint16) error {
	return fmt.Errorf("Item out of stock")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

//main

func main() {
	vendingMachine := New(1, 20) //1 machine of prize 20
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.insertMoney(20)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
