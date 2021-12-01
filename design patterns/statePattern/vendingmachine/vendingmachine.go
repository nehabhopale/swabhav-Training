package vendingmachine

import (
	"fmt"
	"machine/hasitem"
	"machine/hasmoney"
	"machine/itemreq"
	"machine/noitem"
	//"machine/state"
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
	hasItemState := &hasitem.HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemreq.ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasmoney.HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noitem.NoItem{
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
