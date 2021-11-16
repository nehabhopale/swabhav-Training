package main

import (
	customer "bank/bankCustomer"
	"fmt"
)

type MyError struct{
	errMsg string
}
type (er MyError)Error() string{
	return er.errMsg
}
func err (err error) {
	neha := customer.New("saving",customer.bankCustomer{1, "neha", 10000})
	pooja := customer.New("savin",customer.bankCustomer{3, "pooja", 2000})
	// fmt.Println("neha's actual account balance -", neha.GetBalance())
	// fmt.Println("pooja'sactual account balacne -", pooja.GetBalance())
	if pooja.GetAccountType()!="saving"{
		err =MyError{"Sorry we can transfer only from saving account"}
		return 
	}
	validTransfer := neha.TransferMoney(pooja, 1000)
	if validTransfer {
		fmt.Println("neha's account balance after deduction-", neha.GetBalance())
		fmt.Println("pooja's account balacne after deduction-", pooja.GetBalance())
	} else {
		fmt.Println("Sorry,your account balance is  not sufficient")
	}
	//validTransfer := neha.TransferMoney(pooja, -1000)
	//fmt.Println(pooja.GetBalance())

}
