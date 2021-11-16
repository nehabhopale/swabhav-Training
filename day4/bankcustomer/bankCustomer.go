package bankCustomer

type bankCustomer struct {
	bankId  int
	name    string
	balance uint32
}

func (c *bankCustomer) GetName() string {
	return c.name
}
func (c *bankCustomer) SetName(nameOfCustomer string) {
	c.name = nameOfCustomer
}
func (c *bankCustomer) SetBalance(newBalance uint32) {
	c.balance = newBalance
}
func (c *bankCustomer) GetBalance() uint32 {
	return c.balance
}

func (c *bankCustomer) TransferMoney(toOtherCustomer *bankCustomer, amountToBeTransfered uint32) bool {
	if c.balance >= amountToBeTransfered {
		c.balance = c.balance - amountToBeTransfered
		toOtherCustomer.balance = toOtherCustomer.balance + amountToBeTransfered
		return true

	}
	return false
}

func New(bankId int, name string, balance uint32) *bankCustomer {
	return &bankCustomer{
		bankId:  bankId,
		name:    name,
		balance: balance,
	}
}
