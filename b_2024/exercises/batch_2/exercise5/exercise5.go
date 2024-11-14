package main

import (
	"fmt"
	"sort"
)

/*
	Exercise 5: Unique Elements in a Slice

	Task: Write a program that finds the unique elements in a slice of integers.

		Create a function uniqueElements that takes a slice of integers and returns a new slice with only the unique elements.
		In the main function, define a slice of integers with some duplicate values.
		Call the uniqueElements function and print the resulting slice of unique elements.
*/

func uniqueElements(slice []int) []int {
	// Based on this solution
	// https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}

	for _, elem := range slice {
		if !uniqueMap[elem] {
			uniqueMap[elem] = true
			uniqueSlice = append(uniqueSlice, elem)
		}
	}
	// Additionally sorted it
	sort.Ints(uniqueSlice)
	return uniqueSlice
}

func main() {
	fmt.Println("Batch 2 - Exercise 5")
	myslice := []int{1, 6, 4, 3, 5, 8, 2, 6, 7, 7, 9}
	fmt.Println(myslice)
	result := uniqueElements(myslice)
	fmt.Println("Unique elements:", result)

}
