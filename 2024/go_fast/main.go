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
	// key:value pairs
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint{"Adam": 23, "Ahsoka": 5}
	fmt.Println(myMap2["Ahsoka"])
	// if asking for a non existent value -- default is 0
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}
	// remove a value from a map
	delete(myMap2, "Adam")

	//Loops
	// Iterate over maps, arrays or slices - range
	for name := range myMap2 {
		age = myMap2[name]
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	// "While loop"
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for qa := 0; qa < 10; qa++ {
		fmt.Println(qa)
	}

	// Strings, Runes, Bytes
	var myString = "resume"
	// rune --> int32 encoded alias
	// runes are immutable
	var indexed = myString[0] // Adds variable to an array
	// it stores the value of the string
	fmt.Println(myString)
	fmt.Println(indexed)                     // returns 114
	fmt.Printf("%v, %T\n", indexed, indexed) // value, type

	// UTF-8
	// strings - binary numbers (ascii)
	// Own conversion pattern for non ascii characters

	// Structs and interfaces
	// Structs - Collection of fields
	// SEE OUTSIDE OF MAIN FUNCTION
	type myStruct0 struct {
		X         int
		Y         int
		Z         int
		int             // only defining a type by its variable type
		ownerInfo owner //outside of main function
	}

	//fmt.Println(myStruct0{1, 2, 3, 4})
	fmt.Println(owner{"Owner"})
	/*
		v := myStruct0{1, 2}
		v.X = 4
		fmt.Println(v.X)
	*/
	var mynum myStruct0 = myStruct0{X: 100, Y: 200, ownerInfo: owner{"Alex"}} //asigning values to multiple fields at a time
	mynum.Z = 300
	mynum.int = 2 //single field value assignment
	fmt.Println(mynum.X, mynum.Y, mynum.Z, mynum.int, mynum.ownerInfo.name)

	// Struct methods
	// Functions directly related to a struct
	var myEngine gasEngine = gasEngine{25, 15} // "Class" is created/instanceated
	fmt.Println(myEngine.mpg, myEngine.litres)
	fmt.Printf("Total miles: %v", myEngine.milesleft()) //Calling method/"class"

	// interfaces
	// pending to understand

	// pointers
	// store memory locations instead of values
	// var p *int32 //nil
	// var ir int32 // 0
	// & and *
	x := 7
	fmt.Println(&x) // where it is stored

}

//	================
//
// NON MAIN FUNCTIONS
//
//	================
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

// structs

type owner struct {
	name string
}
type gasEngine struct {
	mpg    uint8
	litres uint8
}

func (e gasEngine) milesleft() uint8 {
	// (e gasEngine) assigns the function to the type
	return e.litres * e.mpg
}
