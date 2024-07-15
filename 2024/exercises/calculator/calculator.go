package main

import (
	"fmt"
)

func add(a, b int) int {

	return a + b
}

func extract(c, d int) int {

	return c - d
}

func multiply(e, f int) int {

	return e * f
}

func divide(g, h int) int {

	return g / h
}

func controlf(j int) string {

	if j%2 == 0 {
		// Sidenote: Difference between / and % in Golang
		// / Just makes the division
		// % Gets the remainder (resto, in spanish)
		return "Even"
	} else {
		return "Odd"
	}

}

func main() {

	// Welcome to my Golang manual basic calculator!

	fmt.Println("Sum ", add(1, 9))
	fmt.Println("Minus", extract(10, 9))
	fmt.Println("Multiply", multiply(2, 2))
	fmt.Println("Divide", divide(40, 2))

	fmt.Println("It is ", controlf(10))

}
