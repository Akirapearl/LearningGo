package main

import "fmt"

func isLeapYear(year int) bool {
	/*
			###Exercise 2: Functions and Control Structures

			Task: Write a program that determines whether a given year is a leap year.

		    Create a function that takes an integer year as an argument and returns a boolean. -- Check

		    In the function, use control structures to determine if the year is a leap year:

		    A year is a leap year if it is divisible by 4. -- Check

		    However, if the year is divisible by 100, it is not a leap year unless it is also divisible by 400.-- Check

		    In the main function, prompt the user to enter a year, call isLeapYear, and print whether the year is a leap year.
	*/

	if year%4 == 0 {
		return true
	} else if year%100 == 0 && 400 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println("Choose a year:")
	fmt.Println(isLeapYear(2000))

}
