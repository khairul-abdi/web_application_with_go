/*
Exercise HTTP-Client 1: Fetch and Print a Joke

Instruction: Write a Go program that fetches a joke from https://v2.jokeapi.dev/joke/Any?safe-mode and prints the joke setup and delivery to the console. The joke format should handle the JSON structure provided, to understand the JSON structure, you can visit the URL in your browser.

Possible Solution:
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

func main() {
	response, err := http.Get("https://v2.jokeapi.dev/joke/Any?safe-mode")
	if err != nil {
		fmt.Println("Error fetching joke:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var joke JokeResponse
	if err := json.Unmarshal(body, &joke); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	if joke.Error {
		fmt.Println("Failed to fetch a joke")
		return
	}

	fmt.Printf("Setup: %s\nDelivery: %s\n", joke.Setup, joke.Delivery)
}
