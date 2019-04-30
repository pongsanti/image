package route

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var maxMemory int64 = 10 << 20 // 10 MB
var formKey = "file"
var fileDir = "images"

// Create handle file upload
func Create(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(maxMemory)

	file, handler, err := r.FormFile(formKey)
	if err != nil {
		log.Print("Error Retrieving the File")
		log.Print(err)
		return
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("Cannot read upload file ", err)
		return
	}

	filePath := path.Join(path.Dir(fileDir), handler.Filename)
	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		log.Print("Write to temp file error ", err)
		return
	}

	log.Print("Successfully uploaded file")
}
