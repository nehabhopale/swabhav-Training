package main

import (
	"fmt"
	"reflect"
	"test/math"
	"time"
)

func main() {
	fmt.Println("hi")
	fmt.Println("addition of two numbers", math.AddTwoNumber(4, 5))
	num1 := 2
	num2 := 2.4
	name := "neha"
	fmt.Println(reflect.TypeOf(num1))
	fmt.Printf("number2 type is %T\n", num2)
	fmt.Println(reflect.TypeOf(name))
	IfLoop()
	SwitchLoop()
	FoorLoop()
}

func IfLoop() {
	Time := time.Now()
	h := Time.Hour()
	fmt.Println(h)
	if 8 <= h && h < 12 {
		fmt.Println("good morning")
	} else if 12 <= h && h < 16 {
		fmt.Println("afternoon")
	} else if 16 <= h && h <= 20 {
		fmt.Println("evening")
	} else {
		fmt.Println("night")
	}

}

func SwitchLoop() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

}
func FoorLoop() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
		i++
	}
}
