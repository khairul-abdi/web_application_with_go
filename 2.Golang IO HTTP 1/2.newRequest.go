package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/ability/?limit=1", nil)
	if err != nil {
		panic(err)
	}

	// Request dengan method `Do()` pada instance http.Client yang sudah dibuat
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Cetak status code request
	fmt.Println("Status: ", resp.Status)

	// Kita bisa membaca response body menggunakan package ioutil.
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert body type menjadi string dan cetak
	fmt.Println(string(responseData))
}
