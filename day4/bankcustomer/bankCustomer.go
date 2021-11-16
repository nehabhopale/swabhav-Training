package bankCustomer

type bankCustomer struct {
	bankId  int
	name    string
	balance int
}

func (c *bankCustomer) GetName() string {
	return c.name
}
func (c *bankCustomer) SetName(nameOfCustomer string) {
	c.name = nameOfCustomer
}
func (c *bankCustomer) SetBalance(newBalance int) {
	c.balance = newBalance
}
func (c *bankCustomer) GetBalance() int {
	return c.balance
}

func (c *bankCustomer) TransferMoney(toOtherCustomer *bankCustomer, amountToBeTransfered int) bool {
	if c.balance >= amountToBeTransfered {
		c.balance = c.balance - amountToBeTransfered
		toOtherCustomer.balance = toOtherCustomer.balance + amountToBeTransfered
		return true

	}
	return false
}

func New(bankId int, name string, balance int) *bankCustomer {
	return &bankCustomer{
		bankId:  bankId,
		name:    name,
		balance: balance,
	}
}
