package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// membuat koneksi ke server golang.org melalui tcp network
	conn, err := net.Dial("tcp", "golang.org:80")
	// periksa jika ada error saat koneksi ke server
	// akan keluar dari program `log.Fatal(...)`
	if err != nil {
		log.Fatal("The following error occured", err)
	}
	// koneksi tersambung, kirim GET request ke server
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// kirim request "enter" untuk mengakhiri koneksi
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("Send request error ", err)
	}

	fmt.Println("Response: ", status)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		log.Fatal("Failed: ", err)
	}

	defer resp.Body.Close()
	fmt.Println(resp)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
