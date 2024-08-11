package main

import (
	"fmt"
	"os"
)

/*
Exercise 3: Environment Variables

Task: Write a program that reads environment variables and prints them.

	Use the os.Getenv function to read specific environment variables.
	Define a list of common environment variables to read (e.g., HOME, PATH, USER).
	Loop through the list and print each environment variable and its value.
	Handle cases where the environment variable is not set.
*/
func main() {
	fmt.Println("Batch 3 - Exercise 3")

	var EnVar = []string{"HOME", "PATH", "USER"}

	for i := 0; i < len(EnVar); i++ {
		fmt.Println(os.Getenv(EnVar[i]))

		fmt.Println("")
	}
	fmt.Println("Are this all your needed env Variables?[Y/n]")
	var condition string
	fmt.Scanln(&condition)
	if condition == "yes" || condition == "Yes" || condition == "Y" || condition == "y" {
		fmt.Println("Nice! See you next time")
	} else {
		// Ask for input
		fmt.Println("Missing any environment variable? Write it down!")
		// Potential improvement point - Exctract this part so it can be called
		// multiple times
		var w1 string
		fmt.Scanln(&w1)
		// Append variable to array
		EnVar = append(EnVar, w1)
		last := EnVar[len(EnVar)-1]
		//fmt.Println(last)
		fmt.Println(os.Getenv(last))
	}
}
