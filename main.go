// main.go
// test if the build trigger works
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handleHello).Methods("GET")
	fmt.Println("Server is listening on port 8084...")
	http.ListenAndServe(":8084", r)
}
