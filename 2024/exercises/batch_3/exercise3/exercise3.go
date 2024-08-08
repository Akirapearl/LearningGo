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

	var EnVar = [3]string{"HOME", "PATH", "USER"}

	for i := 0; i < len(EnVar); i++ {
		fmt.Println(os.Getenv(EnVar[i]))

		fmt.Println("")
	}
	fmt.Println("Are this all your needed env Variables?[Y/n]")
	if true {
		fmt.Println("Nice! See you next time")
	} else {
		// Ask for input
		// Append into array OR create a new variable to print the desired Env Variable

	}
}
