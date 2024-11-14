package main

import "fmt"

/*
###Exercise 5: Structs

Task: Create a program that stores and displays information about a book.

	Define a struct Book with the following fields: Title, Author, YearPublished, and ISBN.
	Create an instance of the Book struct and initialize it with values.
	Write a function printBookDetails that takes a Book struct as an argument and prints the book's details in a readable format.
	In the main function, call printBookDetails with the book instance.
*/
type Book struct {
	Title      string
	Author     string
	YPublished int
	ISBN       int
}

func main() {
	fmt.Println("Exercise 5")
	var b1 Book
	//var b2 Book

	b1.Title = "Dune"
	b1.Author = "Frank Herbert"
	b1.YPublished = 1965
	b1.ISBN = 987654321

	// Call the function
	printBookDetails(b1)
}

// function that takes a Book struct as an argument
// and prints the book's details in a readable format.
func printBookDetails(b Book) {
	fmt.Printf("Book Name: %s\n", b.Title)
	fmt.Printf("Author Name: %s\n", b.Author)
	fmt.Printf("Year Published %d\n", b.YPublished)
}
