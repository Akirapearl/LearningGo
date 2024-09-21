package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	//w.Write([]byte("Hello from Patatabox"))

	// Use template.ParseFiles() - Read the template file into a template set.
	// If error - log detailed error message - http.Error() - no subsequent code is executed
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// File path passed to template.ParseFiles() - must be either relative to current
	// or an absolute path
	ts, err := template.ParseFiles(files...) //3 dots added automatically

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Then we use Execute() method - write the template content as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
