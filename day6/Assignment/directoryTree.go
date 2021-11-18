package main

import (
	"fmt"
	"io/ioutil"

	//"os"
	"path/filepath"
	//"strings"
)

// func printListing(entry string, depth int) {
// 	//indent := strings.Repeat("|   ", depth)
// 	fmt.Printf("|-- %s\n", entry)
// }

func printDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	//printListing(path, depth)
	for _, entry := range entries {

		for i := 0; i < depth; i++ {
			fmt.Print("___")
		}

		//full_path, _ := os.Readlink(filepath.Join(path, entry.Name()))

		//printListing(entry.Name(), depth+1)
		if entry.IsDir() {
			fmt.Println("{" + entry.Name() + "}")
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			//printListing(entry.Name(), depth+1)
			fmt.Println(entry.Name())
		}
	}
}

func main() {
	path := "C:/Users/neha/Desktop/training/Book"

	printDirectory(path, 0)
}
