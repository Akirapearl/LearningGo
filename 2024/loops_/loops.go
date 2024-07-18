package main

import "fmt"

func Between(a, b int) []int {
	// 8 kyu
	/*
	Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.
	For example:
		a = 1
		b = 4
		--> [1, 2, 3, 4]
	*/
	for i := 0; i < b; i++ {
		a--
		//c = b - a
		var myArray []int
		myArray[i]

	}
	return myArray[]
}
func main() {
	fmt.Println(Between(5, 10))
}
