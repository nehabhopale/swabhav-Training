package main

import (
	customer "bank/bankCustomer"
	"fmt"
)

func main() {
	neha := customer.New(1, "neha", 10000)
	pooja := customer.New(3, "pooja", 2000)
	fmt.Println("neha's actual account balance -", neha.GetBalance())
	fmt.Println("pooja'sactual account balacne -", pooja.GetBalance())

	validTransfer := neha.TransferMoney(pooja, 1000)
	if validTransfer {
		fmt.Println("neha's account balance after deduction-", neha.GetBalance())
		fmt.Println("pooja's account balacne after deduction-", pooja.GetBalance())
	} else {
		fmt.Println("Sorry,your account balance is  not sufficient")
	}

}
