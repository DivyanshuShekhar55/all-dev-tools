package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    myServerMux := http.NewServeMux()
    myServerMux.HandleFunc("/submit", submitFn)
    myServerMux.HandleFunc("/", helloworld)

    server := http.Server{
        Addr:    ":8000",
        Handler: myServerMux,
    }

    fmt.Println("Server starting on :8000")
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

func submitFn(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Submit endpoint")
}

func helloworld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}