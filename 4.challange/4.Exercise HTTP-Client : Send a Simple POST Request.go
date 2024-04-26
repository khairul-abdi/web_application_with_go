/*
Instruction: Write a Go program to send a POST request to your unique webhook URL from Webhook.site. The POST request should send JSON data containing a message field with the value "Hello, webhook!". After sending the request, print the response status code and body. To use Webhook.site, visit the site, copy your unique URL, and use it to inspect your requests.
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://webhook.site/your-unique-url"
	jsonData := []byte(`{"message": "Hello, webhook!"}`)

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

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
