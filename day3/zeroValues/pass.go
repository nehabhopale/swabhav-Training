package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println(num) //10
	fmt.Println(&num)
	increment(num)
	fmt.Println(num) //10 becaz increement num incremeneted its inner value
	fmt.Println(&num)
	refrenceIncrement(&num) //increment the above val
	fmt.Println(num)
	//var q *int
	//fmt.Println(*q)
}
func increment(num int) {
	fmt.Println("inside increment func add of num", &num)
	num = num + 20 //this increement only the value num whose scope is inside num only
}
func refrenceIncrement(num *int) { // num variable of type pointer pointing to int type eg.& num is pointer pointing to int type
	fmt.Println("inside ref function value of passed num", num)
	*num = *num + 30 //this is goint to increment the value present at location num
}
