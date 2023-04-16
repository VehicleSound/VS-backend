package controller

import (
	"context"
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
	"mime/multipart"
)

type FileUseCase interface {
	UploadImage(file *multipart.FileHeader) (string, error)
	UploadSound(file *multipart.FileHeader) (string, error)
	GetSound(id string) (*domain.File, error)
	GetImage(id string) (*domain.File, error)
}

type FileController struct {
	u FileUseCase
}

func NewFile(u FileUseCase) *FileController {
	return &FileController{u: u}
}

func (c *FileController) GetImage(ctx context.Context, req *dto.GetFileRequest) (*dto.GetFileResponse, error) {
	file, err := c.u.GetImage(req.FileId)
	if err != nil {
		return nil, err
	}

	return &dto.GetFileResponse{
		MimeType: "image/jpeg",
		Bytes:    file.Bytes,
		Size:     file.Size,
	}, nil
}

func (c *FileController) GetSound(ctx context.Context, req *dto.GetFileRequest) (*dto.GetFileResponse, error) {
	file, err := c.u.GetSound(req.FileId)
	if err != nil {
		return nil, err
	}

	return &dto.GetFileResponse{
		MimeType: "audio/mpeg",
		Bytes:    file.Bytes,
		Size:     file.Size,
	}, nil
}

func (c *FileController) UploadImage(ctx context.Context, req *dto.UploadFileRequest) (*dto.UploadFileResponse, error) {
	id, err := c.u.UploadImage(req.File)
	if err != nil {
		return nil, err
	}

	return &dto.UploadFileResponse{FileId: id}, nil
}

func (c *FileController) UploadSound(ctx context.Context, req *dto.UploadFileRequest) (*dto.UploadFileResponse, error) {
	id, err := c.u.UploadSound(req.File)
	if err != nil {
		return nil, err
	}

	return &dto.UploadFileResponse{FileId: id}, nil
}
