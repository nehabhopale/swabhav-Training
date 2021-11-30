package main

import "fmt"

type account struct {
	accNo uint64
}

type securityCode struct {
	code uint16
}
type wallet struct {
	balance int
}

func (a account) CheckAccount(accNo uint64) bool {
	return a.accNo == accNo

}
func (c securityCode) CheckCode(code uint16) bool {
	return c.code == code

}
func (w *wallet) GetBalance() int {
	return w.balance
}

func (w *wallet) Credit(amount int) int {
	w.balance = w.balance + amount
	return w.balance
}
func (w *wallet) Debit(amount int) int {
	w.balance = w.balance - amount
	return w.balance
}

type ledger struct {
	newEntry []string
}

func (l *ledger) makeEntry(entry string) {
	l.newEntry = append(l.newEntry, entry)
}

type notification struct {
	msg string
}

func (n *notification) sendNotification(msg string) {
	n.msg = msg
	fmt.Println(msg)
}

func NewAcc(accNo uint64) *account {
	return &account{
		accNo: accNo,
	}
}
func NewSecure(code uint16) *securityCode {
	return &securityCode{
		code: code,
	}
}
func NewWallet(balance int) *wallet {
	return &wallet{
		balance: balance,
	}
}
func main() {

	neha := NewAcc(23456789)
	code := NewSecure(4234)
	wallet := NewWallet(500)

	if (neha.CheckAccount(23456789)) && (code.CheckCode(4234)) {
		wallet.Credit(500)
		fmt.Println(wallet.GetBalance())
		wallet.Debit(200)
		fmt.Println(wallet.GetBalance())
		var entry = make([]string, 10)
		bankLedger := ledger{
			newEntry: entry,
		}
		notification := notification{
			msg: "",
		}
		bankLedger.makeEntry("money transfer of 500 to  neha")
		notification.sendNotification("money transfer of 500 to neha")

	}

}
