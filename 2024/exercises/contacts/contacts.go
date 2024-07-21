package main

import (
	"fmt"
	"math/rand"
)

/*
   Exercise: Simple Contact List

       Task: Create a program that allows the user to add and view contacts.
	   Each contact should have a name and a phone number.
       Use a map to store the contacts, where the key is the contact's name
       and the value is the phone number.

	STATUS: ONGOING
*/

func main() {
	fmt.Println("Contacts exercise")

	// initialize map
	contacts := make(map[string]int)

	// Create a "default" value
	num := rand.Intn(1000) // Generate a random number between 0 and 99

	contacts["Alex"] = num

	fmt.Println(contacts["Alex"])

	// Reads input & generates variable for the input requested
	fmt.Println("Input name:")
	var w1 string
	fmt.Scanln(&w1)

	// Assigns a number to it && prints it
	contacts[w1] = num
	fmt.Println("read line: ", contacts[w1])

}
