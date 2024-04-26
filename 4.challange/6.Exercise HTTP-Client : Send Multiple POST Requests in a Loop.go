/*
Instruction: Write a Go program to send five POST requests in a loop to your webhook URL. Each POST request should send JSON data with a number field incrementing from 1 to 5. Print the response status and body for each request.
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	url := "https://webhook.site/your-unique-url"

	for i := 1; i <= 5; i++ {
		jsonData := []byte(`{"number": ` + strconv.Itoa(i) + `}`)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Printf("Request %d: Status - %s, Body - %s\n", i, resp.Status, string(body))
	}
}
