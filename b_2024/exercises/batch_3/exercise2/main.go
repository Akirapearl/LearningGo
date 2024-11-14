package main

import (
	"fmt"
	"net/http"
)

/*
Exercise 2: Simple HTTP Server

Task: Write a simple HTTP server that responds with "Hello, World!" to any request.

    Create a new HTTP server in the main function.
    Define a handler function that writes "Hello, World!" to the response.
    Use the http.HandleFunc function to register the handler for the root path.
    Start the HTTP server on port 8080 and handle any errors.

*/

func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	message := fmt.Sprintf("You requested: %s", path)

	fmt.Printf("%v", message)
}

/*
    printf(const char *format, ...) is used to print the data onto the standard output which is often a computer monitor.
    sprintf(char *str, const char *format, ...) is like printf. Instead of displaying the formated string on the standard output i.e. a monitor, it stores the formated data in a string pointed to by the char pointer (the very first parameter). The string location is the only difference between printf and sprint syntax.
    fprintf(FILE *stream, const char *format, ...) is like printf again. Here, instead of displaying the data on the monitor, or saving it in some string, the formatted data is saved on a file which is pointed to by the file pointer which is used as the first parameter to fprintf. The file pointer is the only addition to the syntax of printf.

	If stdout file is used as the first parameter in fprintf, its working is then considered equivalent to that of printf.
	https://progressivecoder.com/how-to-create-a-simple-http-web-server-in-go-a-step-by-step-approach/

*/

func main() {

	http.HandleFunc("/", handleRequest)

	fmt.Println("Hello world! I hear you from port 8080!")

	http.ListenAndServe(":8080", nil)
}
