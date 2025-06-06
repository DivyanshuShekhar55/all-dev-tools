package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	myServerMux := http.NewServeMux()
	myServerMux.HandleFunc("/submit", submitFn)
	myServerMux.HandleFunc("/", helloworld)

	server := http.Server{
		Addr:    "8000",
		Handler: myServerMux,
	}

}

func submitFn(w http.ResponseWriter, r *http.Request) {
	/* TODO */
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	/* TODO */
}
