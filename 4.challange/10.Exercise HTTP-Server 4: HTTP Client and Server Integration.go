/*
Instruction: Create an HTTP server that fetches a joke from the Joke API (https://v2.jokeapi.dev/joke/Any) whenever it receives a GET request on /joke. If the joke is deemed inappropriate (any flags are true), the server should respond with "Inappropriate joke detected", otherwise, it should display the joke setup and delivery.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JokeResponse struct {
	Error    bool   `json:"error"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
}

func fetchJoke() (JokeResponse, error) {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		return JokeResponse{}, fmt.Errorf("error fetching joke: %v", err)
	}
	defer resp.Body.Close()

	var joke JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return JokeResponse{}, fmt.Errorf("error decoding joke: %v", err)
	}
	return joke, nil
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	joke, err := fetchJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if joke.Flags.NSFW || joke.Flags.Religious || joke.Flags.Political || joke.Flags.Racist || joke.Flags.Sexist || joke.Flags.Explicit {
		fmt.Fprintln(w, "Inappropriate joke detected")
	} else {
		fmt.Fprintf(w, "Setup: %s\nDelivery: %s\n", joke.Setup, joke.Delivery)
	}
}

func main() {
	http.HandleFunc("/joke", jokeHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
