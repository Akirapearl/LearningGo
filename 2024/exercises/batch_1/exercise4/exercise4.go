package main

import (
	"fmt"
	"strings"
)

/*
###Exercise 4: Strings, Runes, and Bytes

Task: Write a program that counts the number of vowels in a given string.

    Prompt the user to enter a string.
    Convert the string to lowercase to handle case insensitivity.
    Iterate through the string, character by character (rune by rune).
    Count the vowels (a, e, i, o, u).
    Print the total number of vowels found.
*/

func main() {
	fmt.Println("Exercise 4")

	// Ask for input
	fmt.Println("Input name:")
	var w1 string
	fmt.Scanln(&w1)

	w2 := strings.ToLower(w1)
	sum := 0

	// Convert the string to a slice of runes to handle multibyte characters
	runes := []rune(w2)

	// Iterate over the slice of runes using a for loop
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\n", runes[i])

		if runes[i] == 'a' || runes[i] == 'e' || runes[i] == 'i' {
			sum++
		} else if runes[i] == 'o' || runes[i] == 'u' {
			sum++
		}
	}

	fmt.Println("Total de vocales: ", sum)
}
