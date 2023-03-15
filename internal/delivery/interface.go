package delivery

import (
	dto2 "github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
)

type AuthController interface {
	SignIn(req *dto2.AuthRequest) (*dto2.AuthResponse, error)
	ValidateToken(token string) (*dto2.TokenResponse, error)
}
type UserController interface {
	Register(req *dto2.RegisterRequest) (*dto2.RegisterResponse, error)
	ChangeLogin(req *dto2.ChangeLoginRequest) error
	ChangeEmail(req *dto2.ChangeEmailRequest) error
	ChangePassword(req *dto2.ChangePasswordRequest) error
	GetUser(req *dto2.GetUserRequest) (*dto2.GetUserResponse, error)
	GetUserById(id string) (*dto2.GetUserResponse, error)
	AddToFav(req *dto2.AddToFavRequest) error
}
type TagController interface {
	CreateTag(req *dto2.CreateTagRequest) (*dto2.CreateTagResponse, error)
	GetAllTags() ([]*dto2.TagResponse, error)
	GetTagById(id string) (*dto2.TagResponse, error)
}

type FileController interface {
	UploadImage(req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error)
	UploadSound(req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error)
}
type SearchController interface {
	Search(req *dto2.SearchRequest) ([]*dto2.SoundResponse, error)
}
type SoundController interface {
	GetAllSounds() ([]*dto2.SoundResponse, error)
	GetSoundById(id string) (*dto2.SoundResponse, error)
	CreateSound(t *dto2.TokenResponse, req *dto2.CreateSoundRequest) (*dto2.CreateSoundResponse, error)
	GetRandomSounds(limit int) ([]*dto2.SoundResponse, error)
}
