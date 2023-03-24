package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	name    string
	address string
}

var userdetails []User

func FormHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	files := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/", files.ServeHTTP)
	// x-www-form-urlencoded
	http.HandleFunc("/createuser", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		userdetails = append(userdetails, User{
			name:    name,
			address: address,
		})
		fmt.Println(userdetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User Created"))
	})
	// raw data
	http.HandleFunc("/createraw", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Read Data!"))
	})
	// multipart/form-data
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			fmt.Println(err)
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("File Uploaded!"))
	})
	http.ListenAndServe(":8000", nil)
}
