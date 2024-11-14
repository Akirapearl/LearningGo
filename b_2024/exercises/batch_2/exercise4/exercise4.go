package main

import (
	"fmt"
	"sort"
)

/*
Exercise 4: Sorting Algorithms

Task: Implement a simple sorting algorithm (e.g., Bubble Sort) to sort a list of integers.

	Create a function bubbleSort that takes a slice of integers and sorts it in ascending order.
	In the main function, define a slice of integers and print it.
	Call bubbleSort to sort the slice.
	Print the sorted slice.
*/

func bubbleSort(slice []int) {
	sort.Ints(slice)

}
func main() {
	fmt.Println("Batch 2 - Exercise 4")
	myslice := []int{1, 6, 4, 3, 5, 2}
	fmt.Println(len(myslice))
	bubbleSort(myslice)
	fmt.Println(myslice)
}
