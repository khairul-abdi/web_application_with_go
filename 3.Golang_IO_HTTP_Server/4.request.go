package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		// Read request Method
		method := r.Method

		// Read URL and parameter value:
		url := r.URL
		urlParam := r.URL.Query().Get("param")

		// Read request Header
		headerValues := r.Header.Values("foo") // menampilkan ["bar1", "bar2"]
		headerGet := r.Header.Get("foo")       // menampilkan "bar1" . Hanya nilai pertama yang dikembalikan.

		// Read request Body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Error read Body" + err.Error()))
			return
		}

		writer := fmt.Sprintf("Method:\t\t %v \nUrl:\t\t %v \nUrl param:\t %v \nHeaderValues:\t %v \nHeaderGet:\t %v \nBody:\t\t %v", method, url, urlParam, headerValues, headerGet, string(body))
		w.Write([]byte(writer))
	})

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
