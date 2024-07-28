package main

import (
	"fmt"
	"strings"
)

func HoopCount(n int) string {
	// 8 kyu
	/*
			Alex just got a new hula hoop, he loves it but feels discouraged because his little brother is better than him.

			Write a program where Alex can input (n) how many times the hoop goes round and it will return him an encouraging message:

		    If Alex gets 10 or more hoops, return the string "Great, now move on to tricks".
		    If he doesn't get 10 hoops, return the string "Keep at it until you get it".
	*/
	if n >= 10 {
		return "Great, now move on to tricks"
	} else if n < 10 {
		return "Keep at it until you get it"
	} else {
		return "Value not accepted"
	}
}

func Initials(name string) string {
	// 8 kyu
	/*
		Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

		The output should be two capital letters with a dot separating them.

		It should look like this:

			Sam Harris => S.H

			patrick feeney => P.F
	*/
	parts := strings.Split(name, " ")
	if len(parts) != 2 {
		return ""
	}

	firstInitial := strings.ToUpper(string(parts[0][0]))
	secondInitial := strings.ToUpper(string(parts[1][0]))

	return fmt.Sprintf("%s.%s", firstInitial, secondInitial)
}

func BreakChocolate(n, m int) int {
	/*
		Your task is to split the chocolate bar of given dimension n x m into small squares. Each square is of size 1x1 and unbreakable.

		Implement a function that will return minimum number of breaks needed.
		For example if you are given a chocolate bar of size 2 x 1 you can split it to single squares in just one break,
		but for size 3 x 1 you must do two breaks.

		If input data is invalid you should return 0 (as in no breaks are needed if we do not have any chocolate to split).
		Input will always be a non-negative integer.

	*/

	return 0
}
func main() {
	fmt.Println(HoopCount(10))
	fmt.Println(Initials("Hugh Jackman"))
	fmt.Println(Initials("anna paquin"))
}
