package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*
Exercise 6: CSV File Handling

Task: Write a program that reads a CSV file, processes the data, and prints the results.

    Read CSV File:
        Prompt the user to enter the filename of a CSV file.
        Use the encoding/csv package to read the file.
        Assume the CSV file contains a list of users with columns Name, Age, and Email.

    Print Data:
        Print the details of each user.
*/

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("Batch 3 - Exercise 1")

	// Open the CSV file
	file, err := os.Open("data.csv")
	check(err)
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err2 := reader.ReadAll()
	check(err2)

	for _, row := range data {
		name := row[0]
		age := row[1]
		mail := row[2]
		fmt.Printf("%v %v %v\n", name, age, mail)

	}
}
