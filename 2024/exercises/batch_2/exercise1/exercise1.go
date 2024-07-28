package main

import "fmt"

/*
Exercise 1: Factorial Calculation

Task: Write a program that calculates the factorial of a given number.

    Create a function factorial that takes an integer n and returns the factorial of n using recursion.
    In the main function, prompt the user to enter a number.
    Call the factorial function with the input number and print the result.
*/

func factorial(number int) int {
	var result int = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println("Batch 2 - Exercise 1")

	var num int
	fmt.Println("Input name:")
	fmt.Scanln(&num)

	fmt.Printf("Factorial %d", factorial(num))
}

/*
for i:= 1;i <= number;i++{
	result *= uint64(i)

}*/
