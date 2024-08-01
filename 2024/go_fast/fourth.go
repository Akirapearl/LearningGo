package main

import "fmt"

/*
import "fmt"

func printNumber(item, defaultValue int) (int, int) {

	return item, defaultValue
}

func printString(item, defaultValue string) (string, string) {

	return item, defaultValue
}

func main() {
	fmt.Println("Generics")

	num1, num2 := printNumber(10, 20)
	num3, num4 := printString("one", "two")

	fmt.Printf("%v \n%v\n", num1, num2)
	fmt.Printf("%v \n%v", num3, num4)

}
*/

// first generic
func printItem[T any](item, defaultValue T) (T, T) {

	return item, defaultValue
}

func main() {
	fmt.Println("Generics")
	num1, num2 := printItem(1, 2)
	num3, num4 := printItem("tres", "quatre")
	bool1, bool2 := printItem[bool](false, true)
	fmt.Printf("%v \n%v\n", num1, num2)
	fmt.Println(num3, num4)
	fmt.Println(bool1, bool2)
}
