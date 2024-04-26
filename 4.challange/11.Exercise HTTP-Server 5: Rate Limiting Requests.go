/*
Instruction: Implement an HTTP server that applies rate limiting to its endpoints. Use a simple algorithm to allow up to 5 requests per minute per IP address for the /rate endpoint. If exceeded, respond with a status code of 429 (Too Many Requests). Use r.RemoteAddr to get the IP address of the client.
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var clients = make(map[string]int)
var mtx sync.Mutex

func rateHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	mtx.Lock()
	count, found := clients[ip]
	if found && count >= 5 {
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		mtx.Unlock()
		return
	}
	if !found || count < 5 {
		clients[ip] = count + 1
		time.AfterFunc(time.Minute, func() {
			mtx.Lock()
			clients[ip] = clients[ip] - 1
			mtx.Unlock()
		})
	}
	mtx.Unlock()
	fmt.Fprintln(w, "Request accepted")
}

func main() {
	http.HandleFunc("/rate", rateHandler)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
