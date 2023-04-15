package file

import (
	"errors"
	"fmt"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/usecase"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type UseCase struct {
	repo  usecase.Repository
	store usecase.Storage
}

func New(repo usecase.Repository, store usecase.Storage) *UseCase {
	return &UseCase{
		repo:  repo,
		store: store,
	}
}

func (u *UseCase) UploadImage(fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	if !u.checkImageExt(ext) {
		return "", errors.New("invalid image extension")
	}

	reader, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}
	defer reader.Close()

	file, err := domain.NewFile(ext, reader, fh.Size)
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	// TODO actions with the image

	if err := u.store.CreateFile("images", file); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	return file.Id, nil
}

func (u *UseCase) UploadSound(fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	if !u.checkSoundExt(ext) {
		return "", errors.New("invalid sound extension")
	}

	reader, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}
	defer reader.Close()

	file, err := domain.NewFile(ext, reader, fh.Size)
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	// TODO actions with the sound

	if err := u.store.CreateFile("sounds", file); err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	return file.Id, nil
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
