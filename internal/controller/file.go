package controller

import (
	"context"
	"github.com/timickb/transport-sound/internal/controller/dto"
	"mime/multipart"
)

type FileUseCase interface {
	UploadImage(file *multipart.FileHeader) (string, error)
	UploadSound(file *multipart.FileHeader) (string, error)
}

type FileController struct {
	u FileUseCase
}

func NewFile(u FileUseCase) *FileController {
	return &FileController{u: u}
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
