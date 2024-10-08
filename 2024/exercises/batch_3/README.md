Exercise 1: CSV File Handling

Task: Write a program that reads a CSV file, processes the data, and prints the results.

    Read CSV File:
        Prompt the user to enter the filename of a CSV file.
        Use the encoding/csv package to read the file.
        Assume the CSV file contains a list of users with columns Name, Age, and Email.

    Print Data:
        Print the details of each user.

Exercise 2: Simple HTTP Server

Task: Write a simple HTTP server that responds with "Hello, World!" to any request.

    Create a new HTTP server in the main function.
    Define a handler function that writes "Hello, World!" to the response.
    Use the http.HandleFunc function to register the handler for the root path.
    Start the HTTP server on port 8080 and handle any errors.

Exercise 3: Environment Variables

Task: Write a program that reads environment variables and prints them.

    Use the os.Getenv function to read specific environment variables.
    Define a list of common environment variables to read (e.g., HOME, PATH, USER).
    Loop through the list and print each environment variable and its value.
    Handle cases where the environment variable is not set.

Exercise 4: Concurrency with Goroutines

Task: Write a program that spawns multiple goroutines to print numbers concurrently.

    Create a function printNumbers that takes an integer id and a channel of integers.
    In the main function, create a channel and launch multiple goroutines that call printNumbers.
    Each goroutine should read numbers from the channel and print them along with its id.
    Send a few numbers to the channel and close it.
    Use a sync.WaitGroup to ensure all goroutines complete before the program exits.