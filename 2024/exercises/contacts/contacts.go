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
*/
func randNumber() int {
	// Create a "default" value -- Generate a random number between 0 and 1000
	num := rand.Intn(1000)
	return num
}

func main() {
	fmt.Println("Contacts exercise")

	// initialize map
	contacts := make(map[string]int)

	contacts["Alex"] = randNumber()

	//fmt.Println(contacts["Alex"])

	// Reads input & generates variable for the input requested
	fmt.Println("Input name:")

	var w1 string
	fmt.Scanln(&w1)
	// Assign number to new element of the map
	// New call to the function - Generates different number
	contacts[w1] = randNumber()

	// Print entire map
	fmt.Println("map:", contacts)

	// Prints assigned number only
	fmt.Println("New contact has number: ", contacts[w1])

	// Remove element from the map, then list content to verify
	delete(contacts, w1)
	fmt.Println("map:", contacts)

}
