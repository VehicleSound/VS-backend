package usecase

import (
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/repository"
)

type Repository interface {
	CreateUser(login, email, pwdHash string) (string, error)
	GetUserByLogin(login string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error)

	CreateTag(title string) (*domain.Tag, error)
	GetTagById(id string) (*domain.Tag, error)
	GetAllTags() ([]*domain.Tag, error)
	GetTagByTitle(title string) (*domain.Tag, error)

	GetAllSounds(limit int) ([]*domain.Sound, error)
	GetSoundsNameLike(name string) ([]*domain.Sound, error)
	GetSoundsByTagId(tagId string) ([]*domain.Sound, error)
	GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error)
}
