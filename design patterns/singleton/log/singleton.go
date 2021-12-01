package main

import (
	"fmt"
)

type log struct {
}

var singleInstance *log

func getInstance() *log {

	if singleInstance == nil {
		fmt.Println("Creting Single Instance Now")
		singleInstance = &log{}
	} else {
		fmt.Println("Single Instance already created-1")
	}

	return singleInstance

}
func main() {

	getInstance()

}
