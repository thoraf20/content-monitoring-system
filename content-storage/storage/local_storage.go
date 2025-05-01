package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(fileHeader *multipart.FileHeader, destDir string) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, 0755)
	}

	dstPath := filepath.Join(destDir, fileHeader.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return dstPath, nil
}
