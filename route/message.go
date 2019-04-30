package route

import (
	"errors"
)

var errConnectingDatabase = errors.New("Problem connecting the database")
var errCannotGetFormFile = errors.New("Cannot get form file")
var errCannotReadFile = errors.New("Cannot read upload file")
var errCannotWriteFile = errors.New("Cannot write file to the storage")

var errImageIDInvalid = errors.New("Image id invalid")
var errImageNotFound = errors.New("Image not found")
var errCannotRemoveFile = errors.New("Cannot remove file")
