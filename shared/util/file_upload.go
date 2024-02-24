package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 32 * 1024 * 1024

type FileUpload struct{}

func NewFileUpload() *FileUpload {
	return &FileUpload{}
}

func (*FileUpload) FileUpload(r *http.Request, pathUpload string) ([]string, error) {

	var arrayfileName []string

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		return nil, err
	}

	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {

		if fileHeader.Size > MAX_UPLOAD_SIZE {
			strError := fmt.Sprintf("The uploaded image is too big: %s .", fileHeader.Filename)
			err := errors.New(strError)
			return nil, err
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}

		defer file.Close()

		//err = os.MkdirAll("../uploads", os.ModePerm)
		err = os.MkdirAll(pathUpload, os.ModePerm)
		if err != nil {
			return nil, err
		}

		id := uuid.New()
		fileName := id.String()
		//f, err := os.Create(fmt.Sprintf("../uploads/%s", fileName))
		f, err := os.Create(pathUpload + fileName)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			return nil, err
		}

		arrayfileName = append(arrayfileName, fileName)
	}

	return arrayfileName, nil
}
