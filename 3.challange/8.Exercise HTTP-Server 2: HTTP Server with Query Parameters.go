/*
Instruction: Extend the basic HTTP server to handle a new route /greet. This route should accept a name query parameter and respond with "Hello, [name]!" If no name is given, it should respond with "Hello, Guest!"
*/

package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet", greetHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
