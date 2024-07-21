package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Exercise: Parse JSON

Task: Write a program that reads a JSON string representing a list
of people (with fields for name and age) and parses it into a struct. --> before it was an slice of structs
Then, print out the names and ages of all people in the list.

STATUS: PENDING
*/
func main() {
	fmt.Println("This is a JSON parser")
	// Read entire file at once
	content, err := os.ReadFile("read.json")

	if err != nil {
		// Throw error if the file is NULL
		log.Fatal(err)
	}
	// Print content
	fmt.Println(string(content))

	// Go over the document word by word
	file, err2 := os.Open("read.json")

	if err2 != nil {
		log.Fatal(err2)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	// Parse JSON into a struct

	// JSON decoder
	parser := json.NewDecoder(file)
	var person people

	err = parser.Decode(&person)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %s", err)
	}

	// reflect --> It enables you to examine the type and value of variables dynamically.

	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)
	fmt.Printf("City: %s\n", person.city)

	file.Close()
}

// Struct declaration
type people struct {
	name string
	age  uint8
	city string
}
