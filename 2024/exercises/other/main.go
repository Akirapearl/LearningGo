package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// This file references to "Learn Go Fast: Full Tutorial video"
	// https://www.youtube.com/watch?v=8uiZC0l4Ajw

	textmessage := "Hello World"
	fmt.Println(len(textmessage)) // gets number of bites
	// docs: https://www.includehelp.com/golang/len-function-with-examples.aspx
	fmt.Println(utf8.RuneCountInString(textmessage)) // actual length

	var myRune rune = 'a'
	fmt.Println((myRune)) //outputs 97 - expected

}
