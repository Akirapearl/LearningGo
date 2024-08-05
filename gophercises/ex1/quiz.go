package main

/*
https://github.com/gophercises/quiz/tree/master
*/

import (
	"encoding/csv"
	"fmt"
	"os"
)

type question struct {
	Statement string
	Answer    int
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

	correctAnswers := 0
	incorrect := 0
	//count := 0
	for i, row := range data {
		question := row[0]
		answer := row[1]
		fmt.Print("Question ", i+1, " ", question, " ")

		// Listen for input
		var w1 string
		//w1, errif := fmt.Scanln(&w1)
		fmt.Scanf("%v", &w1)
		//check(errif)
		//fmt.Printf("%v", w1)

		if w1 == answer {
			correctAnswers++
		} else {
			incorrect++
		}
	}
	fmt.Println()
	fmt.Println("Total correct answers: ", correctAnswers, "Total incorrects: ", incorrect)
}
