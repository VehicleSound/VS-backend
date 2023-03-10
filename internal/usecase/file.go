package usecase

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileUseCase struct {
	r Repository
}

func NewFileUseCase(r Repository) *FileUseCase {
	return &FileUseCase{r: r}
}

func (u *FileUseCase) UploadImage(savePath string, file *multipart.FileHeader) (string, error) {
	content, err := file.Open()
	if err != nil {
		return "", err
	}
	defer content.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, content); err != nil {
		return "", err
	}

	fileId := uuid.NewString()

	path := filepath.Join(savePath, fileId+".png")
	err = os.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	return fileId, nil
}

func (u *FileUseCase) UploadSound(savePath string, file *multipart.FileHeader) (string, error) {
	// TODO implement
	panic("implement me")
}
