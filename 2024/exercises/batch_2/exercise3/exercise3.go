package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Exercise 3: Basic File I/O

Task: Write a program that reads a text file, counts the number of words in the file, and prints the result.

    Prompt the user to enter the filename.
    Open and read the file.
    Split the text into words and count them.
    Print the total number of words found.
*/

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	fmt.Printf("Batch 2 - Exercise 3\n\n")
	var filename string
	// Ask for input
	fmt.Println("Write a file name stored locally")
	fmt.Scanln(&filename)

	//file, err := os.Open(filename) -- Open file
	//file, err := os.ReadFile(filename) - Directly read file
	file, err := os.ReadFile(filename)

	// If the file does not exist or is declared incorrectly, it launches an error exception
	check(err)
	fts := string(file)

	split := strings.Split(fts, " ")
	countr := 0

	for i := 0; i < len(split); i++ {
		// improvement point, storing values
		// without showing them
		fmt.Println(split[i])
		countr++
	}

	fmt.Println("Total amount of words ", countr)

	//fmt.Printf("%s", split)
	//fmt.Println(reflect.TypeOf(split))

	//fmt.Printf("%s", string(file)) -- Used during ReadFile
	//file.Close()

}
