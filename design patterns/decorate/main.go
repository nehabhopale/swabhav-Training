package main

import ("fmt"
"pizza/pizzamenu")

func main() {

	extravagnzaPizza := &pizzamenu.ExtraVaganza {}

	//Add cheese topping
	extravagnzaPizzaWithCheese := &pizzamenu.CheeseTopping{
		Pizza: extravagnzaPizza,
	}

	fmt.Println("pizza with cheese",extravagnzaPizzaWithCheese.GetPrize())
	//Add tomato topping
	extravagnzaPizzaWithTomato := &pizzamenu.TomatoTopping{
		Pizza: extravagnzaPizza,
	 }
	 fmt.Println("pizza with tomato",extravagnzaPizzaWithTomato.GetPrize())

	//add cheese and tomato toppings
	extravagnzaPizzaWithCheeseAnTomato:=&pizzamenu.TomatoTopping{
		Pizza: extravagnzaPizzaWithCheese,
	}
	fmt.Println("pizza with cheese and tomato",extravagnzaPizzaWithCheeseAnTomato.GetPrize())
	
}