package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a file server - ./ui/static directory
	// The path given to the http.Dir function is relative
	// to the project directory root

	fileServer := http.FileServer(http.Dir("./ui/static"))

	// mux.Handle() - register the file server as handler for
	// all URL paths that startart with /static/

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
