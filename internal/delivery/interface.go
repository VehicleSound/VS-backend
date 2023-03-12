package delivery

import (
	"github.com/timickb/transport-sound/internal/controller/dto"
)

type AuthController interface {
	SignIn(req *dto.AuthRequest) (*dto.AuthResponse, error)
	ValidateToken(token string) (*dto.TokenResponse, error)
}
type UserController interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ChangeLogin(req *dto.ChangeLoginRequest) error
	ChangeEmail(req *dto.ChangeEmailRequest) error
	ChangePassword(req *dto.ChangePasswordRequest) error
	GetUser(req *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetUserById(id string) (*dto.GetUserResponse, error)
}
type TagController interface {
	CreateTag(req *dto.CreateTagRequest) (*dto.CreateTagResponse, error)
	GetAllTags() ([]*dto.TagResponse, error)
	GetTagById(id string) (*dto.TagResponse, error)
}

type FileController interface {
	UploadImage(req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
	UploadSound(req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
}
type SearchController interface {
	Search(req *dto.SearchRequest) ([]*dto.SoundResponse, error)
}
type SoundController interface {
	GetAllSounds() ([]*dto.SoundResponse, error)
	GetSoundById(id string) (*dto.SoundResponse, error)
	CreateSound(t *dto.TokenResponse, req *dto.CreateSoundRequest) (*dto.CreateSoundResponse, error)
	GetRandomSounds(limit int) ([]*dto.SoundResponse, error)
}
