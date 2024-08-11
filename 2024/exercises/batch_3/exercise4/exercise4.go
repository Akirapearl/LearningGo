package main

import "fmt"

/*
Exercise 4: Concurrency with Goroutines

Task: Write a program that spawns multiple goroutines to print numbers concurrently.

	Create a function printNumbers that takes an integer id and a channel of integers.
	In the main function, create a channel and launch multiple goroutines that call printNumbers.
	Each goroutine should read numbers from the channel and print them along with its id.
	Send a few numbers to the channel and close it.
	Use a sync.WaitGroup to ensure all goroutines complete before the program exits.
*/

func printNumbers() {

}
func main() {
	fmt.Println("Batch 3 - Exercise 4")

}
