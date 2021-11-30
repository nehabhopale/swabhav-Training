package main

import (
	acc "acc/account"
)

func main() {
	neha := acc.NewAcc(23456789)
	code := acc.NewSecure(4234)
	balance := acc.NewWallet(500)

	if (neha.CheckAccount(23456789)) && (code.CheckCode(4234)) {

	}

}
