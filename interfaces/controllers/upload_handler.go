package controllers

import (
	"os"
	"net/http"
	"time"
	"path/filepath"
	"io"

//	"github.com/ucpr/go-echo-clean-arch/domain"
//	"../../usecase"
)

func UploadHandler(c Context) (err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		t := time.Now()
		filename := t.Format("2006-01-02_150405")
		ext := filepath.Ext(file.Filename)

		dst, err := os.Create("~/"+ filename + ext)
		if err != nil {
			return err
		}

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, "Uploaded successfully")
}

