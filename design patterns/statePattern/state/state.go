package state

type State interface {
	addItem(uint8) error
	requestItem() error
	insertMoney(money uint16) error
	dispenseItem() error
}
