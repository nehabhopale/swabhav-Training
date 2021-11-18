package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Print("Enter your filepath: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	filepath := scanner.Text()
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// get the size
	size := fi.Size()
	
	fmt.Print("Enter the size format you want: ")
	scaning := bufio.NewScanner(os.Stdin)
	scaning.Scan()
	sizeFormat := scaning.Text()
	size.printSize(sizeFormat)
}
func (b ByteSize) printSize( sizeFormat string) string {
	// fi, err := os.Stat(filepath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // get the size
	// size := fi.Size()
	switch{
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	}
}
