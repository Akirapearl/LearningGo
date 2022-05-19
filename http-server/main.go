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
// info: https://www.reddit.com/r/golang/comments/usye7v/comment/i96jtx7/?utm_source=share&utm_medium=web2x&context=3
    http.ServeFile(w, r, "./files/index.html")
}

func main() {

    http.HandleFunc("/hello", hello)
    //http.HandleFunc("/headers", headers) -- Calls
    http.Handle("/", http.FileServer(http.Dir("./files")))

    http.ListenAndServe(":8090", nil)
}
