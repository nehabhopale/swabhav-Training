package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	val := "writing using ioutil\n"
	data := []byte(val)

	err1 := ioutil.WriteFile("data.txt", data, 0)

	if err1 != nil {
		log.Fatal(err)
	}

	fmt.Println("done")

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("second line"); err != nil {
		log.Fatal(err)
	}

	//Print the contents of the file
	data1, err22 := ioutil.ReadFile("data.txt")
	if err22 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data1))
}
