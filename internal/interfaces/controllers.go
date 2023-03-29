package interfaces

import (
	"context"
	dto2 "github.com/timickb/transport-sound/internal/controller/dto"
)

type AuthController interface {
	SignIn(ctx context.Context, req *dto2.AuthRequest) (*dto2.AuthResponse, error)
	GetUserByToken(ctx context.Context, token string) (*dto2.TokenResponse, error)
}
type UserController interface {
	Register(ctx context.Context, req *dto2.RegisterRequest) (*dto2.RegisterResponse, error)
	ChangeLogin(ctx context.Context, req *dto2.ChangeLoginRequest) error
	ChangeEmail(ctx context.Context, req *dto2.ChangeEmailRequest) error
	ChangePassword(ctx context.Context, req *dto2.ChangePasswordRequest) error
	GetUser(ctx context.Context, req *dto2.GetUserRequest) (*dto2.GetUserResponse, error)
	GetUserById(ctx context.Context, id string) (*dto2.GetUserResponse, error)
	AddToFav(ctx context.Context, req *dto2.AddToFavRequest) error
}
type TagController interface {
	CreateTag(ctx context.Context, req *dto2.CreateTagRequest) (*dto2.CreateTagResponse, error)
	GetAllTags(ctx context.Context) ([]*dto2.TagResponse, error)
	GetTagById(ctx context.Context, id string) (*dto2.TagResponse, error)
}

type FileController interface {
	UploadImage(ctx context.Context, req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error)
	UploadSound(ctx context.Context, req *dto2.UploadFileRequest) (*dto2.UploadFileResponse, error)
}
type SearchController interface {
	Search(ctx context.Context, req *dto2.SearchRequest) ([]*dto2.SoundResponse, error)
}
type SoundController interface {
	GetAllSounds(ctx context.Context) ([]*dto2.SoundResponse, error)
	GetSoundById(ctx context.Context, id string) (*dto2.SoundResponse, error)
	CreateSound(ctx context.Context, t *dto2.TokenResponse, req *dto2.CreateSoundRequest) (*dto2.CreateSoundResponse, error)
	GetRandomSounds(ctx context.Context, limit int) ([]*dto2.SoundResponse, error)
}
