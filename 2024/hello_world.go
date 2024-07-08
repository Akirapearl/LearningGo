package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, Go!")

	// This is a comment for my first Golang script

	/*
		This is a longer comment. Right?
			fmt.Println("Test")
	*/

	// Variables --> var name type

	var myString string = "Esto es un string"
	// Sidenote: Always double quotes for STRINGS

	fmt.Println(myString)

	myString = "This is a string"
	fmt.Println(myString)

	var mysecString string = "this is the second one"
	// Declared warning as content for the variable was not declared
	//var mysecString string
	//mysecString = "this is the second one"
	fmt.Println(mysecString)

	// strong typing for variables --> can't be changed from string to int or any other type

	var myInt int = 7
	fmt.Println(myInt)
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)

	//Concatenate
	fmt.Println(myString, myInt)

	// Check type of a variable
	fmt.Println(reflect.TypeOf(myInt))

}
