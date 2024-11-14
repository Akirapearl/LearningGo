package main

import (
	"fmt"
	"math"
)

const b float64 = 3.14159

func areaC(a int) float64 {
	/*
			###Exercise 1: Constants, Variables, and Basic Data Types

				Task: Create a program that calculates the area and circumference of a circle.

		    	Define a constant for Pi (π = 3.14159). Declare a variable for the radius of the circle (e.g., radius).
		    	Calculate the area using the formula: Area = π * radius^2. -- Check
	*/

	radius := float64(a)
	pi := b

	square := math.Pow(radius, 2)
	formula := pi * square

	return formula
}

func circm(c int) float64 {
	/*
		Calculate the circumference using the formula: Circumference = 2 * π * radius. -- Check
	*/
	radius := float64(c)
	pi := b

	return 2 * pi * radius
}
func main() {

	fmt.Println(areaC(4))
	fmt.Println(circm(9))

}
