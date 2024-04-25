/*
Instruction: Create an HTTP server that has a single route, /task, which behaves differently based on the HTTP method used (GET to fetch a task description, POST to create a task, and DELETE to delete a task). Respond with appropriate messages for each method:

GET: "Fetching task..."
POST: "Creating a new task..."
DELETE: "Deleting task..."
*/

package main

import (
	"fmt"
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Fetching task...")
	case "POST":
		fmt.Fprintln(w, "Creating a new task...")
	case "DELETE":
		fmt.Fprintln(w, "Deleting task...")
	default:
		http.Error(w, "Unsupported request method.", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/task", taskHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
