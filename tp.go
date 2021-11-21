package main

import "fmt"

func main() {
	type cust struct {
		name string
		id   int
	}

	x := cust{"neha", 1}
	y := cust{"kli", 2}
	var c []cust
	c = append(c, x)
	fmt.Println(c)
}
