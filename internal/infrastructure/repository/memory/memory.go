package memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
	"strings"
	"time"
)

type soundTag struct {
	soundId string
	tagId   string
}

type favourite struct {
	userId  string
	soundId string
}

type Repository struct {
	users      map[string]*domain.User
	tags       map[string]*domain.Tag
	sounds     map[string]*domain.Sound
	soundTags  []*soundTag
	files      map[string]string
	favourites []*favourite
}

func NewRepository() *Repository {
	return &Repository{
		users:      make(map[string]*domain.User),
		tags:       make(map[string]*domain.Tag),
		sounds:     make(map[string]*domain.Sound),
		files:      make(map[string]string),
		soundTags:  make([]*soundTag, 0),
		favourites: make([]*favourite, 0),
	}
}

func (m Repository) AddTagToSound(soundId, tagId string) error {
	if _, ok := m.sounds[soundId]; !ok {
		return errors.New("sound not found")
	}
	if _, ok := m.tags[tagId]; !ok {
		return errors.New("tag not found")
	}

	m.soundTags = append(m.soundTags, &soundTag{
		soundId: soundId,
		tagId:   tagId,
	})
	return nil
}

func (m Repository) CreateSound(sound *domain.Sound) error {
	m.sounds[sound.Id] = sound
	return nil
}

func (m Repository) CreateFile(id, ext string) error {
	m.files[id] = ext
	return nil
}

func (m Repository) GetFileExtById(id string) (string, error) {
	ext, ok := m.files[id]
	if !ok {
		return "", errors.New("file not found")
	}
	return ext, nil
}

func (m Repository) AddFavourite(userId, soundId string) error {
	if _, ok := m.users[userId]; !ok {
		return errors.New("user not found")
	}
	if _, ok := m.sounds[soundId]; !ok {
		return errors.New("sound not found")
	}

	m.favourites = append(m.favourites, &favourite{
		userId:  userId,
		soundId: soundId,
	})

	return nil
}

func (m Repository) CreateUser(login, email, pwdHash string) (string, error) {
	for _, user := range m.users {
		if user.Login == login {
			return "", errors.New("login already exists")
		}
		if user.Email == email {
			return "", errors.New("email already exists")
		}
	}

	user := &domain.User{
		Id:           uuid.NewString(),
		Login:        login,
		Email:        email,
		PasswordHash: pwdHash,
		Confirmed:    false,
		Active:       true,
		DateCreated:  time.Now(),
	}

	m.users[user.Id] = user
	return user.Id, nil
}

func (m Repository) GetUserByLogin(login string) (*domain.User, error) {
	for _, user := range m.users {
		if user.Login == login {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m Repository) GetUserByEmail(email string) (*domain.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m Repository) GetUserById(id string) (*domain.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (m Repository) EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error) {
	if _, ok := m.users[id]; !ok {
		return nil, errors.New("user not found")
	}
	if payload.Email != "" {
		m.users[id].Email = payload.Email
	}
	if payload.Login != "" {
		m.users[id].Login = payload.Login
	}
	if payload.Password != "" {
		m.users[id].PasswordHash = payload.Password
	}

	return m.users[id], nil
}

func (m Repository) CreateTag(title string) (*domain.Tag, error) {
	tag := &domain.Tag{
		Id:    uuid.NewString(),
		Title: title,
	}
	m.tags[tag.Id] = tag
	return tag, nil
}

func (m Repository) GetTagById(id string) (*domain.Tag, error) {
	tag, ok := m.tags[id]
	if !ok {
		return nil, errors.New("tag not found")
	}
	return tag, nil
}

func (m Repository) GetAllTags() ([]*domain.Tag, error) {
	tags := make([]*domain.Tag, len(m.tags))
	for _, tag := range m.tags {
		tags = append(tags, tag)
	}
	return tags, nil
}

func (m Repository) GetTagByTitle(title string) (*domain.Tag, error) {
	for _, tag := range m.tags {
		if tag.Title == title {
			return tag, nil
		}
	}
	return nil, errors.New("tag not found")
}

func (m Repository) GetTagsForSound(soundId string) ([]*domain.Tag, error) {
	tagIds := make([]string, 0)
	for _, item := range m.soundTags {
		if item.soundId == soundId {
			tagIds = append(tagIds, item.tagId)
		}
	}
	tags := make([]*domain.Tag, 0)

	for _, id := range tagIds {
		tag, ok := m.tags[id]
		if !ok {
			return nil, errors.New(fmt.Sprintf("tag with id %s doesn't exist", id))
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (m Repository) GetSoundById(id string) (*domain.Sound, error) {
	sound, ok := m.sounds[id]
	if !ok {
		return nil, errors.New("sound not found")
	}

	return sound, nil
}

func (m Repository) GetAllSounds() ([]*domain.Sound, error) {
	sounds := make([]*domain.Sound, 0)
	for _, value := range m.sounds {
		sounds = append(sounds, value)
	}
	return sounds, nil
}

func (m Repository) GetSounds(limit, offset int) ([]*domain.Sound, error) {
	sounds := make([]*domain.Sound, 0)

	if offset > len(m.sounds) {
		return nil, errors.New("offset exceeds rows count")
	}

	currOffset := 0
	currLimit := 0
	for _, value := range m.sounds {
		if currOffset >= offset && currLimit <= limit {
			sounds = append(sounds, value)
		}
		currLimit++
		currOffset++
	}

	return sounds, nil
}

func (m Repository) GetSoundsNameLike(name string) ([]*domain.Sound, error) {
	sounds := make([]*domain.Sound, 0)
	for _, value := range m.sounds {
		if strings.Contains(value.Name, name) {
			sounds = append(sounds, value)
		}
	}
	return sounds, nil
}

func (m Repository) GetSoundsByTagId(tagId string) ([]*domain.Sound, error) {
	sounds := make([]*domain.Sound, 0)

	for _, sound := range m.sounds {
		for _, tag := range sound.Tags {
			if tag.Id == tagId {
				sounds = append(sounds, sound)
			}
		}
	}

	return sounds, nil
}

func (m Repository) GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error) {
	sounds := make([]*domain.Sound, 0)

	for _, sound := range m.sounds {
		if sound.VehicleId == vehicleId {
			sounds = append(sounds, sound)
		}
	}

	return sounds, nil
}
