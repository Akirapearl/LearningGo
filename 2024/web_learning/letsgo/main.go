package main

// page 27/404 Let's Go by Alex Edwards.
import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Patatabox"))
}

// Second handler function - Non-home result
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

// Third handler function - Non-home result
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form...hehe"))

}

func main() {
	// Initialize router -- servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)                         // HandleFunc assumes already the pass of a ResponseWritter and a Request
	mux.HandleFunc("/snippet/view", snippetView)         // Calls to non home/root results
	mux.HandleFunc("/snippet/create/{$}", snippetCreate) // Dollar means match a single slash, followed by nothing else

	// hostname matching
	// It is possible to include hostnames in route patterns, useful
	// if using Go to redirect HTTP request to multiple sites or services
	/*
		mux := http.NewServeMux()
			mux.HandleFunc("foo.example.org/", fooHandler)
			mux.HandleFunc("bar.example.org/", barHandler)
			mux.HandleFunc("/baz", bazHandler)
	*/

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
