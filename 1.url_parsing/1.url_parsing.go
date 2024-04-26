package main

import (
	"fmt"
	"net/url"
)

// // var urlString = "https://www.google.com/search?q=golang&hl=id"
// var urlString = "https://neurons.ruangguru.com/course/backend-development/web-application/golang-i-o/url-parsing"

// func main() {
// 	var uri, err = url.Parse(urlString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Schema: ", uri.Scheme)     // mendapatkan informasi schema dari URL
// 	fmt.Println("Host: ", uri.Host)         // mendapatkan informasi hostname dari URL
// 	fmt.Println("Path: ", uri.Path)         // mendapatkan informasi path dari URL
// 	fmt.Println("RawQuery: ", uri.RawQuery) // mendapatkan informasi query dari URL
// }

var urlString = "https://www.google.com/search?q=golang&hl=id"

// var urlString = "https://neurons.ruangguru.com/course/backend-development/web-application/golang-i-o/url-parsing"

func main() {
	var uri, err = url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Schema: ", uri.Scheme)     // mendapatkan informasi schema dari URL
	fmt.Println("Host: ", uri.Host)         // mendapatkan informasi hostname dari URL
	fmt.Println("Path: ", uri.Path)         // mendapatkan informasi path dari URL
	fmt.Println("RawQuery: ", uri.RawQuery) // mendapatkan informasi query dari URL
}
