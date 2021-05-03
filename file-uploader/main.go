package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	// Parse input - type multipart/form-data
	r.ParseMultipartForm(5 << 10)

	// Retrieve file from posted form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded file: %v\n", handler.Filename)
	fmt.Printf("File size: %v\n", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)

	// Write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	// Return success/failure status of file upload attempt
	fmt.Fprintf(w, "Successfully uploaded file\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("Golang File Upload Tutorial")
	setupRoutes()
}
