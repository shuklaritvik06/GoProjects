package controllers

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuklaritvik06/GoProjects/filecompress/backend/utils"
)

func UploadFile(r *http.Request) string {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Println(err)
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "Error Retrieving the File"
	}
	defer file.Close()
	if result, _ := utils.Exists("./uploads"); result == false {
		os.Mkdir("./uploads", 0755)
	}
	tempFile, err := ioutil.TempFile("./uploads", "*"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	return tempFile.Name()
}

func Compress(w http.ResponseWriter, r *http.Request) {
	name := UploadFile(r)
	os.Chdir("./uploads")
	data, _ := os.ReadFile(strings.Split(name, "/")[2])
	newfile := strings.Split(strings.Split(name, "/")[2], ".")[1] + ".gz"
	os.Chdir("..")
	if result, _ := utils.Exists("./compressed"); result == false {
		os.Mkdir("./compressed", 0755)
	}
	compress, _ := ioutil.TempFile("./compressed", "*"+newfile)
	p := gzip.NewWriter(compress)
	p.Write(data)
	p.Close()
	file, _ := filepath.Abs("./compressed/" + strings.Split(compress.Name(), "/")[2])
	compressed_data, _ := os.ReadFile(file)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(compressed_data)
}

func Decompress(w http.ResponseWriter, r *http.Request) {
	name := UploadFile(r)
	os.Chdir("./uploads")
	file, _ := os.Open(strings.Split(name, "/")[2])
	defer file.Close()
	reader, _ := gzip.NewReader(file)
	defer reader.Close()
	newfile := strings.Split(strings.Split(name, "/")[2], ".")[1]
	os.Chdir("..")
	if result, _ := utils.Exists("./decompressed"); result == false {
		os.Mkdir("./decompressed", 0755)
	}
	decompressed_file, _ := ioutil.TempFile("./decompressed", "*"+newfile)
	defer decompressed_file.Close()
	_, err := io.Copy(decompressed_file, reader)
	if err != nil {
		fmt.Println(err)
	}
	abspath, _ := filepath.Abs("./decompressed/" + strings.Split(decompressed_file.Name(), "/")[2])
	absfiledata, _ := os.ReadFile(abspath)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(absfiledata)
}
