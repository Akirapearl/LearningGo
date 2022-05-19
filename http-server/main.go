//https://gobyexample.com/http-servers
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello there!\n")
}

func headers(w http.ResponseWriter, r *http.Request) {

    http.ServeFile(w, r, "./files/index.html")
}

func main() {

    http.HandleFunc("/hello", hello)
    //http.HandleFunc("/headers", headers) -- Calls
    http.Handle("/", http.FileServer(http.Dir("./files")))

    http.ListenAndServe(":8090", nil)
}