package interfaces

import (
	"context"
	"github.com/timickb/transport-sound/internal/controller/dto"
)

type AuthController interface {
	SignIn(ctx context.Context, req *dto.AuthRequest) (*dto.AuthResponse, error)
	GetUserByToken(ctx context.Context, token string) (*dto.TokenResponse, error)
}
type UserController interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ChangeLogin(ctx context.Context, req *dto.ChangeLoginRequest) error
	ChangeEmail(ctx context.Context, req *dto.ChangeEmailRequest) error
	ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) error
	GetUser(ctx context.Context, req *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetUserById(ctx context.Context, id string) (*dto.GetUserResponse, error)
	AddToFav(ctx context.Context, req *dto.AddToFavRequest) error
}
type TagController interface {
	CreateTag(ctx context.Context, req *dto.CreateTagRequest) (*dto.CreateTagResponse, error)
	GetAllTags(ctx context.Context) ([]*dto.TagResponse, error)
	GetTagById(ctx context.Context, id string) (*dto.TagResponse, error)
}

type FileController interface {
	UploadImage(ctx context.Context, req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
	UploadSound(ctx context.Context, req *dto.UploadFileRequest) (*dto.UploadFileResponse, error)
}
type SearchController interface {
	Search(ctx context.Context, req *dto.SearchRequest) ([]*dto.SoundResponse, error)
}
type SoundController interface {
	GetAllSounds(ctx context.Context) ([]*dto.SoundResponse, error)
	GetSoundById(ctx context.Context, id string) (*dto.SoundResponse, error)
	CreateSound(ctx context.Context, t *dto.TokenResponse, req *dto.CreateSoundRequest) (*dto.CreateSoundResponse, error)
	GetRandomSounds(ctx context.Context, limit int) ([]*dto.SoundResponse, error)
}
