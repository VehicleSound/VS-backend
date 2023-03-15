package usecase

import (
	domain2 "github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
	"time"
)

type Repository interface {
	CreateUser(login, email, pwdHash string) (string, error)
	GetUserByLogin(login string) (*domain2.User, error)
	GetUserByEmail(email string) (*domain2.User, error)
	GetUserById(id string) (*domain2.User, error)
	EditUser(id string, payload *repository.UserEditPayload) (*domain2.User, error)

	CreateTag(title string) (*domain2.Tag, error)
	GetTagById(id string) (*domain2.Tag, error)
	GetAllTags() ([]*domain2.Tag, error)
	GetTagByTitle(title string) (*domain2.Tag, error)
	GetTagsForSound(soundId string) ([]*domain2.Tag, error)

	AddTagToSound(soundId, tagId string) error
	CreateSound(sound *domain2.Sound) error
	GetSoundById(id string) (*domain2.Sound, error)
	GetAllSounds() ([]*domain2.Sound, error)
	GetSounds(limit, offset int) ([]*domain2.Sound, error)
	GetSoundsNameLike(name string) ([]*domain2.Sound, error)
	GetSoundsByTagId(tagId string) ([]*domain2.Sound, error)
	GetSoundsByVehicleId(vehicleId string) ([]*domain2.Sound, error)

	CreateFile(id, ext string) error
	GetFileExtById(id string) (string, error)

	AddFavourite(userId, soundId string) error
}
type UserContext interface {
	User() *domain2.User
	CreatedAt() time.Time
	Get(key string) interface{}
	Add(key string, val interface{})
}
