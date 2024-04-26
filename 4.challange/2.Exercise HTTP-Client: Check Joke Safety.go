/*
Instruction: Enhance the previous program to check the safety of the joke by examining the flags field. Print whether the joke is safe for all audiences based on the flags (i.e., all flags should be false for the joke to be considered "safe"). If the joke is safe, print "The joke is safe for all audiences", otherwise, print "The joke is not safe for all audiences".
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
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
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

	safe := !joke.Flags.NSFW && !joke.Flags.Religious && !joke.Flags.Political &&
		!joke.Flags.Racist && !joke.Flags.Sexist && !joke.Flags.Explicit

	if safe {
		fmt.Println("The joke is safe for all audiences")
	} else {
		fmt.Println("The joke is not safe for all audiences")
	}
}
