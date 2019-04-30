package route

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pongsanti/image/db/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const ImageIDUrlParam = "imageID"

func DeleteImageHandlerFunc(db *sql.DB, config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Image: CreateDeleteImageHandlerFunc")

		renderError := func(err error) {
			render.JSON(w, r, struct{ Error string }{
				err.Error(),
			})
		}

		ctx := r.Context()
		imageIDString := chi.URLParam(r, ImageIDUrlParam)

		imgID, err := strconv.Atoi(imageIDString)
		if err != nil {
			renderError(errImageIDInvalid)
			return
		}

		// new tx
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Print("Error creating tx ", err)
			renderError(errConnectingDatabase)
			return
		}

		// find image
		img, err := models.FindImage(ctx, tx, imgID)
		if err != nil {
			log.Print("Error find image ", err)
			tx.Rollback()
			renderError(errImageNotFound)
			return
		}

		// delete file
		filePath := filepath.Join(config.FileDir, img.Href)
		log.Printf("Deleting: %s\n", filePath)
		err = os.Remove(filePath)
		if err != nil {
			log.Print("Error remove the physical file ", err)
			tx.Rollback()
			renderError(errCannotRemoveFile)
			return
		}

		// delete record
		count, err := img.Delete(ctx, tx)
		if err != nil {
			log.Print("Error delete an image ", err)
			tx.Rollback()
			renderError(err)
			return
		}
		log.Printf("%d row deleted\n", count)

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			renderError(errConnectingDatabase)
			return
		}

		render.JSON(w, r, img)
	}
}
