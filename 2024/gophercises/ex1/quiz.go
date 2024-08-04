package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type question struct {
	question string
	answer   string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Welcome to the tty-quizz conquest")

	// Open the CSV file
	file, err := os.Open("problems.csv")
	check(err)
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err2 := reader.ReadAll()
	check(err2)

	fmt.Println(data[0][0])
	/*
		// Print the CSV data
		for _, row := range data {
			for _, col := range row {
				fmt.Printf("%s,", col)
			}
			fmt.Println()
		}*/
}
