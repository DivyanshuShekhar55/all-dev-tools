package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

// define the body of the incoming request
type req struct {
	Name string `json:"name"`
}

func main() {

	// start a connection pool to pg server
	ctx := context.Background()
	pool, err := connect()

	if err != nil {
		log.Fatal("unable to connect", err)
		return
	}
	//fmt.Println(pool)
	defer pool.Close()

	// add the user schema if not exists
	initSchema(pool, ctx)

	// create a mux router
	myServerMux := http.NewServeMux()
	myServerMux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		submitFn(w, r, pool, ctx)
	})
	myServerMux.HandleFunc("/", helloworld)

	// create and run the server
	server := http.Server{
		Addr:    ":8000",
		Handler: myServerMux,
	}

	fmt.Println("Server starting on :8000")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func submitFn(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool, ctx context.Context) {
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
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received Name is %s, sending to db ...", payload.Name)

	insertName(payload.Name, pool, ctx)

}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
