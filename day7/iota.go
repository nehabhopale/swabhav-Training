package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter your filepath: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	filepath := scanner.Text()
	fmt.Print("Enter the size format you want: ")
	scaning := bufio.NewScanner(os.Stdin)
	scaning.Scan()
	sizeFormat := scaning.Text()
	printSize(filepath, sizeFormat)
}
func printSize(filepath string, sizeFormat string) {
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// get the size
	size := fi.Size()
	fmt.Println(size)
}
