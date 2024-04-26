/*
Instruction: Modify the previous exercise to send a PUT request instead of a POST request. Add custom headers: X-Custom-Header set to MyValue and send the same JSON data. Print the response status code and body.
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

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Custom-Header", "MyValue")

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
