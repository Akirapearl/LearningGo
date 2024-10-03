package main

// requires to run as go run ./cmd/web (from root of the project)
import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a command-line flag with default value ":4000" w/some help text
	// Value will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse command-line flag
	// Reads value and assings it to addr variable
	// otherwise - always contain default value :4000
	// if errror - application will be terminated
	flag.Parse()

	mux := http.NewServeMux()

	// Create a file server - ./ui/static directory
	// The path given to the http.Dir function is relative
	// to the project directory root

	fileServer := http.FileServer(http.Dir("./ui/static"))

	// mux.Handle() - register the file server as handler for
	// all URL paths that startart with /static/

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
