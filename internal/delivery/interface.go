package delivery

import (
	"github.com/timickb/transport-sound/internal/controller"
	"github.com/timickb/transport-sound/internal/domain"
)

type AuthController interface {
	SignIn(req *controller.AuthRequest) (*controller.AuthResponse, error)
}
type UserController interface {
	Register(req *controller.RegisterRequest) (*controller.RegisterResponse, error)
	ChangeLogin(req *controller.ChangeLoginRequest) error
	ChangeEmail(req *controller.ChangeEmailRequest) error
	ChangePassword(req *controller.ChangePasswordRequest) error
}
type TagController interface {
	CreateTag(req *controller.CreateTagRequest) (*controller.CreateTagResponse, error)
	GetAllTags() ([]*controller.TagResponse, error)
	GetTagById(id string) (*controller.TagResponse, error)
}
type SoundController interface {
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
}
type FileController interface {
	UploadImage(req *controller.UploadFileRequest) (*controller.UploadFileResponse, error)
	UploadSound(req *controller.UploadFileRequest) (*controller.UploadFileResponse, error)
}
