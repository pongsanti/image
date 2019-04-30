package route

import (
	"github.com/pongsanti/dbconnect"
	"net/http"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
    fmt.Println("Hello World")
    setupRoutes()
}

func setupRoutes() {
	dbc, _ := dbconnect.NewDBConnect("localhost", "lib", "lib", "lib", "")


    http.HandleFunc("/upload", CreateNewImageHandlerFunc(dbc.Db, DefaultConfig()))
    http.ListenAndServe(":8080", nil)
}
