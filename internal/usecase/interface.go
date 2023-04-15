package usecase

import (
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
	"time"
)

type Repository interface {
	CreateUser(user domain.User) error
	GetUserByLogin(login string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error)

	CreateTag(tag domain.Tag) error
	GetTagById(id string) (*domain.Tag, error)
	GetAllTags() ([]*domain.Tag, error)
	GetTagByTitle(title string) (*domain.Tag, error)
	GetTagsForSound(soundId string) ([]*domain.Tag, error)

	AddTagToSound(soundId, tagId string) error
	CreateSound(sound *domain.Sound) error
	GetSoundById(id string) (*domain.Sound, error)
	GetAllSounds() ([]*domain.Sound, error)
	GetSounds(limit, offset int) ([]*domain.Sound, error)
	GetSoundsNameLike(name string) ([]*domain.Sound, error)
	GetSoundsByTagId(tagId string) ([]*domain.Sound, error)
	GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error)

	CreateFile(id, ext string) error
	GetFileExtById(id string) (string, error)

	AddFavourite(userId, soundId string) error
}
type UserContext interface {
	User() *domain.User
	CreatedAt() time.Time
	Get(key string) interface{}
	Add(key string, val interface{})
}
type Storage interface {
	CreateFile(bucket string, file *domain.File) error
	DeleteFile(bucket, filename string) error
	GetFile(bucket, filename string) (*domain.File, error)
}
