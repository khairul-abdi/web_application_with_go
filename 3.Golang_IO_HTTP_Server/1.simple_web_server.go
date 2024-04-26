package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "how are you today!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
