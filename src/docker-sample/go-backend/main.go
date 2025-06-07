package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type req struct {
	Name string `json:"name"`
}

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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received body: %s\n", string(body))

	var payload req
	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Errorf("error while reading json")
	}

	fmt.Printf("%s", payload.Name)

}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
