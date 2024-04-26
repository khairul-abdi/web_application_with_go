package main

import (
	"fmt"
	"net/http"
)

func handlerRouter1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hallo Router1"))
}

func handlerRouter2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	w.Write([]byte("Hallo Router2"))

}

func main() {
	// Membuat router dengan http.HandleFunc()
	http.HandleFunc("/", handlerRouter1)
	http.HandleFunc("/index", handlerRouter1)
	http.HandleFunc("/data", handlerRouter2)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
