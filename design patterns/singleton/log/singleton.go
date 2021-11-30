package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

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
	wg.Done()
	return singleInstance

}
func main() {

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go getInstance()
	}
	wg.Wait()
}
