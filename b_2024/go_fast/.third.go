package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d \n", i)
	}
}

func sendData(ch chan<- int) {
	// Sending a value into the channel
	ch <- 20
}

func main() {

	// Calling a function (standard)
	//printNumbers()

	// Calling a Goroutine
	// hangs the execution
	go printNumbers()
	fmt.Println("Goroutines")
	time.Sleep(1 * time.Second)
	// Part of the code is executed, then follows the function described as a goroutine

	// Channels
	// ------------------------
	// "Different pieces of your code talk to each other by sending and receiving messages" <-

	// Creating a channel
	ch := make(chan int)

	go sendData(ch)

	// receiving the message and storing it in a variable
	result := <-ch
	fmt.Println(result)

	// Unbuffered channel
	// Capacity is 0
	// each send operation blocks until there is a corresponding receive operation
	// and viceversa
	// DIRECT SYNCHRONOUD COMMUNICATION BETWEEN BOTH ENDS
	unbufferedCh := make(chan int)
	go func() {
		unbufferedCh <- 66
	}()
	// Buffered channels
	// Capacity greater than 0
	// allows multiple values to be sent into the channel without immediate corresponding receiver
	bufferedCh := make(chan string, 2)
	go func() {
		bufferedCh <- "Hello"
		bufferedCh <- "There"
		close(bufferedCh)
	}()

	value := <-unbufferedCh
	fmt.Println("Unbuffered channel: ", value)

	for msg := range bufferedCh {
		fmt.Println("Buffered channel: ", msg)
	}
}
