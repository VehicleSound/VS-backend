package file

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sunshineplan/imgconv"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/converter"
	"github.com/timickb/transport-sound/internal/interfaces"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type UseCase struct {
	r           usecase.Repository
	log         interfaces.Logger
	maxFileSize int
}

func NewFileUseCase(r usecase.Repository, log interfaces.Logger, maxFileSize int) *UseCase {
	return &UseCase{r: r, log: log, maxFileSize: maxFileSize}
}

func (u *UseCase) UploadImage(savePath string, fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	if !u.checkImageExt(ext) {
		return "", errors.New("invalid converter extension")
	}

	file, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	if len(buf.Bytes()) > u.maxFileSize {
		return "", errors.New("file too large")
	}

	image, err := imgconv.Decode(buf)
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	id := uuid.NewString()
	if err := converter.HandleAndSaveImage(savePath, id, image); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	if err := u.r.CreateFile(id, ext); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	return id, nil
}

func (u *UseCase) UploadSound(savePath string, fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if !u.checkSoundExt(ext) {
		return "", errors.New("invalid sound extension")
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

func (u *UseCase) uploadFile(file *multipart.File, path, ext string) (string, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, *file); err != nil {
		return "", err
	}

	if len(buf.Bytes()) > u.maxFileSize {
		return "", errors.New("file too large")
	}

	fileId := uuid.NewString()

	path = filepath.Join(path, fileId+ext)

	err := os.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}

	return fileId, nil
}

func (u *UseCase) checkImageExt(ext string) bool {
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

func (u *UseCase) checkSoundExt(ext string) bool {
	switch ext {
	case ".mp3":
		return true
	case ".ogg":
		return true
	}

	return false
}

func (u *UseCase) checkImageMime(mime string) bool {
	switch mime {
	case "converter/png":
		return true
	case "converter/jpeg":
		return true
	}

	return false
}

func (u *UseCase) checkSoundMime(mime string) bool {
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
