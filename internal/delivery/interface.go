package delivery

import (
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
)

type AuthController interface {
	SignIn(req *dto.AuthRequest) (*dto.AuthResponse, error)
}
type UserController interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ChangeLogin(req *dto.ChangeLoginRequest) error
	ChangeEmail(req *dto.ChangeEmailRequest) error
	ChangePassword(req *dto.ChangePasswordRequest) error
}
type TagController interface {
	CreateTag(req *dto.CreateTagRequest) (*dto.CreateTagResponse, error)
	GetAllTags() ([]*dto.TagResponse, error)
	GetTagById(id string) (*dto.TagResponse, error)
}
type SoundController interface {
	CreateSound(req *dto.CreateSoundRequest) (*dto.CreateSoundResponse, error)
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
}
type FileController interface {
	UploadImage(req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
	UploadSound(req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
}
