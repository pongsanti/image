package route

import (
	"github.com/go-chi/chi"
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

	r := chi.NewRouter()


	r.Post("/upload", CreateNewImageHandlerFunc(dbc.Db, DefaultConfig()))
	r.Delete("/images/{imageID}", DeleteImageHandlerFunc(dbc.Db, DefaultConfig()))
    http.ListenAndServe(":8080", r)
}
