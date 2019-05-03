package image

import (
	"log"
	"github.com/pongsanti/image/db/models"
	"github.com/volatiletech/sqlboiler/boil"
	"context"
)

type imageCreatorContextKey uint

// ImageCreatorCtxKey is the context for image creator id
const ImageCreatorCtxKey imageCreatorContextKey = 0

func beforeInsertHookFunc(ctx context.Context, exec boil.ContextExecutor, image *models.Image) error {
	creatorID, ok := ctx.Value(ImageCreatorCtxKey).(int)
	if ok {
		image.CreatorID = creatorID
	}
	return nil
}

// RegisterBeforeInsertHook registers before insert hook
func RegisterBeforeInsertHook() {
	log.Print("Image:RegisterBeforeInsertHook")
	models.AddImageHook(boil.BeforeInsertHook, beforeInsertHookFunc)
}