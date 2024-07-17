package main

import (
	"fmt"
	"strings"
)

func HoopCount(n int) string {
	// 8 kyu
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
	parts := strings.Split(name, " ")
	if len(parts) != 2 {
		return ""
	}

	firstInitial := strings.ToUpper(string(parts[0][0]))
	secondInitial := strings.ToUpper(string(parts[1][0]))

	return fmt.Sprintf("%s.%s", firstInitial, secondInitial)
}

func main() {
	fmt.Println(HoopCount(10))
	fmt.Println(Initials("Hugh Jackman"))
	fmt.Println(Initials("anna paquin"))
}
