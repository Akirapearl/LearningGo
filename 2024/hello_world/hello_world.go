package main

import (
	"container/list" //Added automatically when typing list.New()
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

	// Strong typing for variables --> can't be changed from string to int or any other type

	var myInt int = 7
	fmt.Println(myInt)
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)

	// Concatenate
	fmt.Println(myString, myInt)

	// Check type of a variable
	fmt.Println(reflect.TypeOf(myInt))

	var myFloat float32 = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(myInt + int(myFloat))     // loses the decimal part
	fmt.Println(myFloat + float32(myInt)) // keeps the decimal part

	var myBool bool = true
	myBool = true
	fmt.Println(myBool)

	// Create automated/shortened variables
	myString3 := "Another string"
	fmt.Println(myString3)

	// Constant variables --> No need to use right away to avoid warnings or errors
	const myConst string = "First constant"
	fmt.Println(myConst)

	// Flux control
	if myInt == 10 {
		fmt.Println("El valor es 10")
	} else if myInt == 11 {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor es diferente a 10")
	}

	myInt = 10
	myString = "Adios"

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10, Hola!")
	} else if myInt == 10 || myString == "Adios" {
		fmt.Println("Uno de los valores es correcto")
	}

	// Arrays

	var myArray [3]int
	myArray[2] = 1
	myArray[1] = 0
	myArray[0] = 2
	// myArray[3] = 8 -> Reports error as Array was defined with only 3 elements
	fmt.Println(myArray)

	var myArray2 [3]string
	myArray2[1] = "test"
	fmt.Println(myArray2)

	// Maps (key:value)

	myMap := make(map[string]int)

	myMap["Test"] = 25
	myMap["Prueba"] = 28
	myMap["Prova"] = 30

	fmt.Println(myMap["Prova"])

	myMap2 := map[string]int{"I'm": 28, "Learning": 24, "Golang": 90}
	fmt.Println(myMap2)

	// List -- Works like a stack

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Loops

	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for index, value := range myMap {
		fmt.Println(index, value)
	}

	// Functions
	fmt.Println(myFunction())

	// Structs

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Test", 25}
	fmt.Println(myStruct)

	fmt.Println(testnumber(5, 5))
}

func myFunction() string {
	return "My function"
	//fmt.Println("this is my function")
}

func testnumber(x, y int) int {
	return x + y
}
