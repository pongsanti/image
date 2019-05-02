package image

import (
	"encoding/json"
	"github.com/pongsanti/image/db/models"
	"time"
)

type ImageRes struct {
	Image *Image
}

type Image models.Image

func (i *Image) MarshalJSON() ([]byte, error) {
	type Alias Image
	return json.Marshal(&struct {
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
		CreatorID *int       `json:"creator_id,omitempty"`
		*Alias
	}{
		Alias:     (*Alias)(i),
		DeletedAt: nil, // override to not marshal
		CreatorID: nil, // override to not marshal
	})
}
