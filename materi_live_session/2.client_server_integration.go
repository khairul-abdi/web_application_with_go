package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
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

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type JokeResponse struct {
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

func fetchJoke() (Joke, error) {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		return Joke{}, fmt.Errorf("error fetching joke: %v", err)
	}
	defer resp.Body.Close()

	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return Joke{}, fmt.Errorf("error decoding joke: %v", err)
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
		respJoke := JokeResponse{
			Setup:    joke.Setup,
			Delivery: joke.Delivery,
		}

		respSuccess := ResponseSuccess{
			Status:  http.StatusOK,
			Message: "berhasil",
			Data:    respJoke,
		}

		resByte, _ := json.Marshal(respSuccess)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resByte)
	}
}

func main() {
	http.HandleFunc("/joke", jokeHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
