package usecase

import (
	domain2 "github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
)

type StubRepo struct{}

func (m StubRepo) CreateUser(login, email, pwdHash string) (string, error) {
	return "", nil
}

func (m StubRepo) GetUserByLogin(login string) (*domain2.User, error) {
	return &domain2.User{}, nil
}

func (m StubRepo) GetUserByEmail(email string) (*domain2.User, error) {
	return &domain2.User{}, nil
}

func (m StubRepo) GetUserById(id string) (*domain2.User, error) {
	return &domain2.User{}, nil
}

func (m StubRepo) EditUser(id string, payload *repository.UserEditPayload) (*domain2.User, error) {
	return &domain2.User{}, nil
}

func (m StubRepo) CreateTag(title string) (*domain2.Tag, error) {
	return &domain2.Tag{}, nil
}

func (m StubRepo) GetTagById(id string) (*domain2.Tag, error) {
	return &domain2.Tag{}, nil
}

func (m StubRepo) GetAllTags() ([]*domain2.Tag, error) {
	return []*domain2.Tag{}, nil
}

func (m StubRepo) GetTagByTitle(title string) (*domain2.Tag, error) {
	return &domain2.Tag{}, nil
}

func (m StubRepo) GetTagsForSound(soundId string) ([]*domain2.Tag, error) {
	return []*domain2.Tag{}, nil
}

func (m StubRepo) GetSoundById(id string) (*domain2.Sound, error) {
	return &domain2.Sound{}, nil
}

func (m StubRepo) GetAllSounds() ([]*domain2.Sound, error) {
	return []*domain2.Sound{}, nil
}

func (m StubRepo) GetSounds(limit, offset int) ([]*domain2.Sound, error) {
	return []*domain2.Sound{}, nil
}

func (m StubRepo) GetSoundsNameLike(name string) ([]*domain2.Sound, error) {
	return []*domain2.Sound{}, nil
}

func (m StubRepo) GetSoundsByTagId(tagId string) ([]*domain2.Sound, error) {
	return []*domain2.Sound{}, nil
}

func (m StubRepo) GetSoundsByVehicleId(vehicleId string) ([]*domain2.Sound, error) {
	return []*domain2.Sound{}, nil
}
