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

func result(prompt int) bool {
	//fmt.Println(prompt)
	//count2 := 1

	// Go over the same CSV but only for even values
	/*
		for _, row := range data {
			for _, col := range row {
				count2++
				if count2%2 == 0 {
					var response question = question{}
					// takes the value for when the column is even
					// and compares it for its odd value
					response.Answer = col
					if prompt == col {
						return true
					} else {
						return false
					}
				}
			}
		}

	*/
	return false
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

	count := 0
	// Print the CSV data
	for _, row := range data {
		for _, col := range row {
			//fmt.Printf("%s,", col)
			count++
			// Get all responses
			if count%2 != 0 {
				// asign struct to current progress of the loop
				// 1st being 5+5
				var ask question = question{}
				ask.Statement = col
				fmt.Printf("Resolve %s\n", ask.Statement)
				// Listen for input
				var w1 int
				//w1, errif := fmt.Scanln(&w1)
				fmt.Scanf("%v", &w1)
				//check(errif)
				//fmt.Printf("%v", w1)

				go result(w1)

			}

		}
	}
}
