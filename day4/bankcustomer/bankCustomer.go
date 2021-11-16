package bankCustomer

type bankCustomer struct {
	bankId  int
	name    string
	balance uint32
}
type account struct {
	acctype string
	details bankCustomer
}

func (a *account) GetAccountType() string {
	return a.acctype
}

func (a *account) GetName() string {
	return a.details.name
}
func (a *account) SetName(nameOfCustomer string) {
	a.details.name = nameOfCustomer
}
func (a *account) SetBalance(newBalance uint32) {
	a.details.balance = newBalance
}
func (a *account) GetBalance() uint32 {
	return a.details.balance
}

func (a *account) TransferMoney(toOtherCustomer *account, amountToBeTransfered uint32) bool {
	if a.details.balance >= amountToBeTransfered {
		a.details.balance = a.details.balance - amountToBeTransfered
		toOtherCustomer.details.balance = toOtherCustomer.details.balance + amountToBeTransfered
		return true

	}
	return false
}

// func New(bankId int, name string, balance uint32) *bankCustomer {
// 	return &bankCustomer{
// 		bankId:  bankId,
// 		name:    name,
// 		balance: balance,
// 	}
// }
func New(acctype string, customerdetails bankCustomer) *account {
	return &account{
		acctype: acctype,
		details: customerdetails,
	}
}
