package controller

import (
	dto2 "github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"mime/multipart"
)

type FileUseCase interface {
	UploadImage(savePath string, file *multipart.FileHeader) (string, error)
	UploadSound(savePath string, file *multipart.FileHeader) (string, error)
}

type FileController struct {
	u FileUseCase
}

func NewFileController(u FileUseCase) *FileController {
	return &FileController{u: u}
}

func (c *FileController) UploadImage(req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error) {
	id, err := c.u.UploadImage("static/images/", req.File)
	if err != nil {
		return nil, err
	}

	return &dto2.UploadFileResponse{FileId: id}, nil
}

func (c *FileController) UploadSound(req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error) {
	id, err := c.u.UploadSound("static/sounds/", req.File)
	if err != nil {
		return nil, err
	}

	return &dto2.UploadFileResponse{FileId: id}, nil
}