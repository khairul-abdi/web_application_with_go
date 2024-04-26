package main

import (
	"fmt"
	"net/http"
	"time"
)

// Di sini terdapat TimeHandler dengan atribut Timezone dan Format
type TimeHandler struct {
	Timezone string
	Format   string
}

// TimeHandler memiliki method ServeHTTP(ResponseWriter, *Request). Method ini mencetak hari, tanggal, bulan, dan tahun ke ResponseWriter.
func (h TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation(h.Timezone)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	t := time.Now().In(loc).Format(h.Format)
	w.Write([]byte(fmt.Sprintf("Today is %v in Timezone %s", t, h.Timezone)))
}

func main() {
	// objek dari TimeHandler dengan nilai Timezone Asia/Jakarta dan Format time.RFC3339
	handler := TimeHandler{
		Timezone: "Asia/Jakarta",
		Format:   time.RFC3339,
	}

	// pada Handle, kita menambahkan rute `/timezone` dan objek yang telah dibuat
	http.Handle("/timezone", handler)

	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
