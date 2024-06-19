package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World-v1"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Server is listening on port 8081...")
	http.ListenAndServe(":8081", nil)
}