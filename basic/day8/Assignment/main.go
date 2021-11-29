package main

import (
	"fmt"
	Customer "hotel/Customer"
	Restaurant "hotel/Restaurant"
	"time"
)

func main() {

	neha := Customer.New("neha", "neha@123", 5612345892, "@124", 1000)
	pooja := Customer.New("pooja", "pooja@345", 7634126739, "@567t", 200)
	shardha := Customer.New("shradha", "shradha@345", 7634187639, "@98t", 300)
	Customer.SignUp(neha)
	Customer.SignUp(pooja)
	Purohit := Restaurant.New("purohit")
	Purohit.AddNewItem("pav bhaji", 80, true, 4)
	Purohit.AddNewItem("tea", 10, false, 4)
	Purohit.AddNewItem("idli", 50, true, 5)
	Purohit.AddNewItem("dosa", 60, true, 4)

	Purohit.GetRestDetails()

	go neha.PlaceOrder(Purohit, "pav bhaji")
	go shardha.PlaceOrder(Purohit, "tea")
	go neha.PlaceOrder(Purohit, "idli")
	go pooja.PlaceOrder(Purohit, "dosa")
	time.Sleep(20 * time.Second)
	fmt.Println(neha.GetBalance())
	fmt.Println(shardha.GetBalance())
	fmt.Println(pooja.GetBalance())
	time.Sleep(4* time.Second)
}
