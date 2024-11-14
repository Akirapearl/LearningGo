package main

import (
	"fmt"
	"sync"
)

/*
Exercise 4: Concurrency with Goroutines

Task: Write a program that spawns multiple goroutines to print numbers concurrently.

    Create a function printNumbers that takes an integer id and a channel of integers.
    In the main function, create a channel and launch multiple goroutines that call printNumbers.
    Each goroutine should read numbers from the channel and print them along with its id.
    Send a few numbers to the channel and close it.
    Use a sync.WaitGroup to ensure all goroutines complete before the program exits.
*/

func printNumbers(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Goroutine %d received: %d\n", id, num)
	}
}

func main() {
	fmt.Println("Batch 3 - Exercise 4")

	var wg sync.WaitGroup

	// Create individual channels for each goroutine
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg.Add(3)

	// Launch goroutines with their respective channels
	go printNumbers(1, ch1, &wg)
	go printNumbers(2, ch2, &wg)
	go printNumbers(3, ch3, &wg)

	// Send numbers to the respective channels
	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	// Close the channels
	close(ch1)
	close(ch2)
	close(ch3)

	wg.Wait()
}
