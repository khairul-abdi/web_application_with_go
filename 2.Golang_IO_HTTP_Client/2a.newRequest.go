package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

type (
	ResponseEcho struct {
		Arguments Arguments `json:"args"`
		Headers   Headers   `json:"headers"`
		Url       string    `json:"url"`
	}

	Arguments struct {
		Foo1 string `json:"foo1"`
		Foo2 string `json:"foo2"`
	}

	Headers struct {
		Host           string `json:"host"`
		XForwardedPort string `json:"x-forwarded-port"`
	}
)

func main() {
	var client = &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://postman-echo.com/get?foo1=bar1&foo2=bar2", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseEcho := ResponseEcho{}
	json.Unmarshal(responseData, &responseEcho)

	fmt.Println(reflect.TypeOf(responseEcho))
	fmt.Println(responseEcho)
}
