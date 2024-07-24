package main

import "fmt"

/*
###Exercise 3: Arrays, Slices, Maps, and Loops

Task: Create a program that calculates the average of a list of numbers.

    Define an array or slice of integers with at least five elements.
    Use a loop to iterate through the elements and calculate the sum.
    Calculate the average by dividing the sum by the number of elements.
    Print the average.
*/

const total int = 15

func main() {
	fmt.Println("Exercise 3")
	// Startup array
	// Variable sets the total number of indexes stored within the array (L. 16)

	var NewArray [total]int
	//NewArray[0] = 1

	for i := 0; i <= total-1; i++ {
		// Define each index for the array at once
		value := i * 2
		NewArray[i] = value
	}
	// Explanation for the (total-1) element instead of the len()
	// since the total value is set to 15, the 0 index for the array
	// counts too, meaning that the highest index would be 14, aka
	// the total set above minus one.

	fmt.Println(len(NewArray))
	//fmt.Println(NewArray[14]) -- Test value assignment
	for i := 0; i < len(NewArray); i++ {
		// Iterate within all its elements
		fmt.Println(NewArray[i])
	}
	// Jump line
	fmt.Printf("\n")
	// Calculate sum of all values
	sum := 0

	for i := 0; i < len(NewArray); i++ {
		sum += NewArray[i]
	}
	fmt.Println("Total is ", sum, "and average is ", sum/total)
}
