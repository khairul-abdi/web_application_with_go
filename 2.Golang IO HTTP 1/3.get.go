package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type headers struct {
	XForwardedProto string `json:"x-forwarded-proto"`
	XForwardedPort  string `json:"x-forwarded-port"`
	Host            string `json:"host"`
}

type JsonResponse struct {
	Headers headers
}

func main() {
	// http.Get Implementation:
	resp, err := http.Get("https://postman-echo.com/get?foo1=bar1&foo2=bar2")
	if err != nil {
		log.Fatal(err)
	}
	// Kita bisa membaca response body menggunakan package ioutil.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Convert body type to string
	sb := string(body)
	log.Printf(sb)

	var jsonResponse JsonResponse

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonResponse)

}
