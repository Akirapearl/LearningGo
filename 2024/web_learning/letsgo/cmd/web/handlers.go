package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// func home converted to a method against *application (main.go struct)
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	//w.Write([]byte("Hello from Patatabox"))

	// Use template.ParseFiles() - Read the template file into a template set.
	// If error - log detailed error message - http.Error() - no subsequent code is executed
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// File path passed to template.ParseFiles() - must be either relative to current
	// or an absolute path
	ts, err := template.ParseFiles(files...) //3 dots added automatically
	if err != nil {
		//log.Print(err.Error())
		// Because the home handler is now a method - it can access its fields,
		// including the structured logger-

		// Creates a log entry at Error level containing the error message
		//app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		app.serverError(w, r, err) // Use the serverError() helper.
		return
	}

	// Then we use Execute() method - write the template content as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//log.Print(err.Error())
		//app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err) // Use the serverError() helper.
	}

}

// Change the signature - defined as a method against *application
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
