package main

// 44 of 404
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Add elements to HEADER
	// MUST be before w.Write
	// Key, Value
	w.Header().Add("Server", "Go")

	//w.Write([]byte("Hello from Patatabox"))
	//io.WriteString(w, "Hello world")
	fmt.Fprint(w, "Hello from Patatabox")
}

// Second handler function - Non-home result
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract value of the id wildcard from the request -- r.PathValue()
	// Convert it to an int -- strconv.Atoi()
	// If can't convert -- error 404
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	/*
		msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
		w.Write([]byte(msg))
	*/
	// Writing reponse body in a more efficient way
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Third handler function - Non-home result
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form...hehe"))

}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Change header result (HTTP Code)
	//w.WriteHeader(201)
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))

}

func main() {
	// Initialize router -- servemux
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // HandleFunc assumes already the pass of a ResponseWritter and a Request
	// Wildcard route patterns
	/*
		Define create more flexible routing rules
			- Only 1 per segment
			- each path segment (the bit between forward slash characters)
			can only contain one wildcard

		The rule for preference between two similar routes (i.e /post/{id} - /post/edit)
		 	- the most specific route pattern wins.
			- Therefore "/post/edit" is the more specific route
			pattern and will take precedent.


		Potential edge case - overlapping route patterns
		i.e "/post/new/{id}" and "/post/{author}/latest" - they both match the
		request path /post/new/latest

			- Go’s servemux considers the patterns to conflict, and will panic at
			runtime when initializing the routes.
			- it’s generally good practice to keep overlaps to a minimum or
			avoid them completely.

	*/
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Calls to non home/root results
	mux.HandleFunc("GET /snippet/create", snippetCreate)  // Dollar means match a single slash, followed by nothing else

	// hostname matching
	// It is possible to include hostnames in route patterns, useful
	// if using Go to redirect HTTP request to multiple sites or services
	/*
		mux := http.NewServeMux()
			mux.HandleFunc("foo.example.org/", fooHandler)
			mux.HandleFunc("bar.example.org/", barHandler)
			mux.HandleFunc("/baz", bazHandler)
	*/
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print log as server is starting
	log.Print("Starting server on port :4000...")

	// Start server
	// Pass two parameters - TCP Network address to listen on
	// and servemux (router). Meaning port as destination and router
	// as "origin"
	err := http.ListenAndServe(":4000", mux)
	// err var is always not nil -- log fatal to record any error happening
	log.Fatal(err)

}
