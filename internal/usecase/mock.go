package usecase

import (
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/repository"
)

type StubRepo struct{}

func (m StubRepo) CreateUser(login, email, pwdHash string) (string, error) {
	return "", nil
}

func (m StubRepo) GetUserByLogin(login string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (m StubRepo) GetUserByEmail(email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (m StubRepo) GetUserById(id string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (m StubRepo) EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error) {
	return &domain.User{}, nil
}

func (m StubRepo) CreateTag(title string) (*domain.Tag, error) {
	return &domain.Tag{}, nil
}

func (m StubRepo) GetTagById(id string) (*domain.Tag, error) {
	return &domain.Tag{}, nil
}

func (m StubRepo) GetAllTags() ([]*domain.Tag, error) {
	return []*domain.Tag{}, nil
}

func (m StubRepo) GetTagByTitle(title string) (*domain.Tag, error) {
	return &domain.Tag{}, nil
}

func (m StubRepo) GetTagsForSound(soundId string) ([]*domain.Tag, error) {
	return []*domain.Tag{}, nil
}

func (m StubRepo) GetSoundById(id string) (*domain.Sound, error) {
	return &domain.Sound{}, nil
}

func (m StubRepo) GetAllSounds() ([]*domain.Sound, error) {
	return []*domain.Sound{}, nil
}

func (m StubRepo) GetSounds(limit, offset int) ([]*domain.Sound, error) {
	return []*domain.Sound{}, nil
}

func (m StubRepo) GetSoundsNameLike(name string) ([]*domain.Sound, error) {
	return []*domain.Sound{}, nil
}

func (m StubRepo) GetSoundsByTagId(tagId string) ([]*domain.Sound, error) {
	return []*domain.Sound{}, nil
}

func (m StubRepo) GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error) {
	return []*domain.Sound{}, nil
}
