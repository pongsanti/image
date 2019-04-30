package image

import (
	"github.com/pongsanti/image/db/models"
	"time"
)

type Image struct {
	ID        int
	CreatedAt time.Time
	Filename  string
	Href      string
}


func NewImage(img *models.Image) Image {
	return Image{
		ID:        img.ID,
		CreatedAt: img.CreatedAt,
		Filename:  img.Filename,
		Href:      img.Href,
	}
}
