package main

import "fmt"

/*
Exercise: FizzBuzz

    Task: Write a program that prints the numbers from 1 to 100.
    But for multiples of three, print "Fizz" instead of the number
    and for the multiples of five, print "Buzz".
    For numbers which are multiples of both three and five, print "FizzBuzz".
*/

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else if i%5 == 0 && i != 5 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
