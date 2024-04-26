package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)
	http.HandleFunc("/view", handleImage)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html")) //html dibaca
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	//mengambil relative path dari proyek
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// membuat nama file baru yang akan disimpan
	filename := handler.Filename
	split := strings.Split(filename, ".")
	fmt.Println("MASUKKK")
	if split[len(split)-1] != "png" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("File tidak di dukung")) // menampilkan image sebagai response
		return
	}
	if alias != "" { //kalau alias disi maka nama file = alias
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	// membentuk lokasi tempat menyimpan file
	fileLocation := filepath.Join(dir, "images", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	// mengisi file baru dengan data dari file yang ter-upload
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("image-name")           // mengambil nama image dari query url
	fileBytes, err := os.ReadFile("./images/" + imageName) // membaca file image menjadi bytes
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("File not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes) // menampilkan image sebagai response
	return
}
