package main

import (
	"fmt"
	"strconv"
)

/*
--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take a look.

As the elves making the final adjustments, they discover that their calibration document (your puzzle input)
has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently,
the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration
value that the Elves now need to recover.

On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/

func main() {
	fmt.Println("Advent of Code - First day")
	var calibration [3]string
	calibration[0] = "1abc2"
	//calibration[1] = "pqr3stu8vwx"
	//var res string
	var count int = 0
	for _, c := range calibration[0] {
		d, err := strconv.Atoi(string(c))
		err = nil
		count++
		if d !=
		fmt.Println(err)

	}
}

/*
	//fmt.Printf(" %c\n", c)
		//value := 0
		//value2 := 0

	if !unicode.IsDigit(c) {
			//d := int(c)
			d := string(c)
			// idea is to replace everything
			res = strings.ReplaceAll(d, "abcdefghijklmnopqrstuvwxyz", "")
			count++
			//value = d
			fmt.Println("value", res)
			//} else if !unicode.IsDigit(c) && count != 1 && count <= 2 {

		}
			-----------------------------------------
	if unicode.IsDigit(c) {
			d := int(c - '0')
			count++

			if count == 2 {
				secnd := d
				fmt.Println("second is", secnd)
			}
			first := d
			fmt.Println("first is", first)
		}
*/
