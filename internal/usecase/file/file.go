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
	ext := strings.ToLower(filepath.Ext(fh.Filename))[1:]

	if !u.checkImageExt(ext) {
		return "", errors.New("invalid image extension")
	}

	reader, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}
	defer reader.Close()

	file, err := domain.NewFile("jpg", reader, fh.Size)
	if err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	exId, err := u.repo.GetFileIdBySum(file.HashString())
	if err == nil {
		// Файл уже существует
		return exId, nil
	}

	if err := u.store.CreateFile("images", file); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	if err := u.repo.CreateFile(file.Id, file.Ext, file.HashString()); err != nil {
		return "", fmt.Errorf("err upload image: %w", err)
	}

	return file.Id, nil
}

func (u *UseCase) UploadSound(fh *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(fh.Filename))[1:]

	if !u.checkSoundExt(ext) {
		return "", errors.New("invalid sound extension")
	}

	reader, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}
	defer reader.Close()

	file, err := domain.NewFile("mp3", reader, fh.Size)
	if err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	exId, err := u.repo.GetFileIdBySum(file.HashString())
	if err == nil {
		// Файл уже существует
		return exId, nil
	}

	if err := u.store.CreateFile("sounds", file); err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	if err := u.repo.CreateFile(file.Id, file.Ext, file.HashString()); err != nil {
		return "", fmt.Errorf("err upload sound: %w", err)
	}

	return file.Id, nil
}

func (u *UseCase) GetSound(id string) (*domain.File, error) {
	file, err := u.store.GetFile("sounds", id+".mp3")
	if err != nil {
		return nil, fmt.Errorf("err get sound: %w", err)
	}

	return file, nil
}

func (u *UseCase) GetImage(id string) (*domain.File, error) {
	file, err := u.store.GetFile("images", id+".jpg")
	if err != nil {
		return nil, fmt.Errorf("err get image: %w", err)
	}

	return file, nil
}

func (u *UseCase) checkImageExt(ext string) bool {
	switch ext {
	case "png":
		return true
	case "jpg":
		return true
	case "jpeg":
		return true
	}

	return false
}

func (u *UseCase) checkSoundExt(ext string) bool {
	switch ext {
	case "mp3":
		return true
	case "ogg":
		return true
	}

	return false
}
