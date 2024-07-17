package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// This file references to "Learn Go Fast: Full Tutorial video"
	// https://www.youtube.com/watch?v=8uiZC0l4Ajw

	textmessage := "Hello World"
	fmt.Println(len(textmessage)) // gets number of bites
	// docs: https://www.includehelp.com/golang/len-function-with-examples.aspx
	fmt.Println(utf8.RuneCountInString(textmessage)) // actual length

	var myRune rune = 'a'
	fmt.Println((myRune)) //outputs 97 - expected

	const pi float32 = 3.1415

	fmt.Println(pi)

	var printValue string = "Hi"
	printMe(printValue)

	var result, remainder int = intDivision(11, 2)
	/*if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result is %v", result)*/
	//} else{
	fmt.Printf("Result is %v, remainder is %v", result, remainder)
	//}

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])

	fmt.Println(&intArr[1]) //memory location

	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7) //Second value points to which element from the slice it is wanted to write to
	fmt.Println(intSlice)
	// Arrays have fixed size, slices adapt to what it is needed to introduce
	fmt.Printf("The capacity of the slice is %v ", cap(intSlice))

	// Append multiple values
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Maps
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	/*
		var err error // nil
			if denominator == 0 {
				err = errors.New("Cannot divide by Zero")
				return 0, 0, err
			}*/
	result := numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
