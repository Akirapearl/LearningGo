package main

import "fmt"

/*
Exercise 2: Palindrome Checker

Task: Write a program that checks if a given string is a palindrome.

    Create a function isPalindrome that takes a string and returns a boolean indicating if the string is a palindrome.
    In the main function, prompt the user to enter a string.
    Call the isPalindrome function with the input string and print whether it is a palindrome.
*/

func isPalindrome(paraule string) (result string, isPal bool) {
	for _, v := range paraule {
		result = string(v) + result
		if result == paraule {
			isPal = true
		} else {
			isPal = false
		}
	}
	return
}

func main() {
	fmt.Println("Batch 2 - Exercise 2")
	var word string
	fmt.Println("Write a word")
	fmt.Scanln(&word)

	fmt.Println(isPalindrome(word))
}
