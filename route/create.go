package route

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/render"
	"github.com/pongsanti/image"
	"github.com/pongsanti/image/db/models"
	"github.com/volatiletech/sqlboiler/boil"
)

const mimeMapContentTypeKey = "Content-Type"

// CreateNewImageHandlerFunc handles file upload
func CreateNewImageHandlerFunc(db *sql.DB, config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Image: CreateNewImageHandlerFunc")

		renderError := func(err error) {
			render.JSON(w, r, struct{ Error string }{
				err.Error(),
			})
		}

		ctx := r.Context()

		r.ParseMultipartForm(config.MaxMemory)

		file, handler, err := r.FormFile(config.FormKey)
		if err != nil {
			log.Print("Error Retrieving the File ", err)
			renderError(errCannotGetFormFile)
			return
		}
		defer file.Close()

		log.Printf("Uploaded File: %+v\n", handler.Filename)
		log.Printf("File Size: %+v\n", handler.Size)
		log.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Print("Cannot read upload file ", err)
			renderError(errCannotReadFile)
			return
		}

		// new tx
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Print("Error creating tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		var fileType string
		contentTypes := handler.Header[mimeMapContentTypeKey]
		if len(contentTypes) > 0 {
			fileType = contentTypes[0]
		}
		// insert to db
		img := &models.Image{
			Filename: handler.Filename,
			FileSize: int(handler.Size),
			FileType: fileType,
		}

		// insert to get id
		err = img.Insert(ctx, tx, boil.Infer())
		if err != nil {
			log.Print("Error insert an image record ", err)
			renderError(errConnectingDatabase)
			tx.Rollback()
			return
		}
		pk := img.ID
		physicalFilename := fmt.Sprintf("%d-%s", pk, handler.Filename)
		filePath := filepath.Join(config.FileDir, physicalFilename)

		// write file to the storage
		err = ioutil.WriteFile(filePath, fileBytes, 0644)
		if err != nil {
			tx.Rollback()
			log.Print("Write to the storage error ", err)
			renderError(errCannotWriteFile)
			return
		}

		// update to id-filename
		img.Href = physicalFilename
		_, err = img.Update(ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			log.Print("Update image file path error ", err)
			renderError(errConnectingDatabase)
			return
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			log.Print("Error commit tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		log.Print("Successfully uploaded file")

		render.JSON(w, r, image.ImageRes{
			Image: (*image.Image)(img),
		})
	}
}
