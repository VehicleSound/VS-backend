package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileUseCase struct {
	r Repository
}

func NewFileUseCase(r Repository) *FileUseCase {
	return &FileUseCase{r: r}
}

func (u *FileUseCase) UploadImage(savePath string, fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	if !u.checkImageExt(ext) {
		return "", errors.New("invalid image extension")
	}

	if !u.checkImageMime(fh.Header.Get("Content-Type")) {
		return "", errors.New("invalid mime type")
	}

	file, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}
	defer file.Close()

	id, err := u.uploadFile(&file, savePath, ext)
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	if err := u.r.CreateFile(id, ext); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	return id, nil
}

func (u *FileUseCase) UploadSound(savePath string, fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if !u.checkSoundExt(ext) {
		return "", errors.New("invalid sound extension")
	}

	if !u.checkSoundMime(fh.Header.Get("Content-Type")) {
		return "", errors.New("invalid mime type")
	}

	file, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}
	defer file.Close()

	id, err := u.uploadFile(&file, savePath, ext)
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	if err := u.r.CreateFile(id, ext); err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	return id, nil
}

func (u *FileUseCase) uploadFile(file *multipart.File, path, ext string) (string, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, *file); err != nil {
		return "", err
	}

	fileId := uuid.NewString()

	path = filepath.Join(path, fileId+ext)

	err := os.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	return fileId, nil
}

func (u *FileUseCase) checkImageExt(ext string) bool {
	switch ext {
	case ".png":
		return true
	case ".jpg":
		return true
	case ".jpeg":
		return true
	}

	return false
}

func (u *FileUseCase) checkSoundExt(ext string) bool {
	switch ext {
	case ".mp3":
		return true
	case ".ogg":
		return true
	}

	return false
}

func (u *FileUseCase) checkImageMime(mime string) bool {
	switch mime {
	case "image/png":
		return true
	case "image/jpeg":
		return true
	}

	return false
}

func (u *FileUseCase) checkSoundMime(mime string) bool {
	switch mime {
	case "audio/aac":
		return true
	case "audio/ogg":
		return true
	case "audio/mpeg":
		return true
	case "audio/wav":
		return true
	}

	return false
}
