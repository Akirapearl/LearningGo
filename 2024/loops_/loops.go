package main

import "fmt"

func main() {
	sum := 1
	// While loop, removing init (i = 0), post (i++) and semicolons
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
