package memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	domain2 "github.com/timickb/transport-sound/internal/domain"
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
	users      map[string]*domain2.User
	tags       map[string]*domain2.Tag
	sounds     map[string]*domain2.Sound
	soundTags  []*soundTag
	files      map[string]string
	favourites []*favourite
}

func NewRepository() *Repository {
	return &Repository{
		users:      make(map[string]*domain2.User),
		tags:       make(map[string]*domain2.Tag),
		sounds:     make(map[string]*domain2.Sound),
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

func (m Repository) CreateSound(sound *domain2.Sound) error {
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

func (m Repository) CreateUser(user domain2.User) error {
	for _, user := range m.users {
		if user.Login == user.Login {
			return errors.New("login already exists")
		}
		if user.Email == user.Email {
			return errors.New("email already exists")
		}
	}

	m.users[user.Id] = &domain2.User{
		Id:           uuid.NewString(),
		Login:        user.Login,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Confirmed:    false,
		Active:       true,
		DateCreated:  time.Now(),
	}

	return nil
}

func (m Repository) GetUserByLogin(login string) (*domain2.User, error) {
	for _, user := range m.users {
		if user.Login == login {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m Repository) GetUserByEmail(email string) (*domain2.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m Repository) GetUserById(id string) (*domain2.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (m Repository) EditUser(id string, payload *repository.UserEditPayload) (*domain2.User, error) {
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

func (m Repository) CreateTag(tag domain2.Tag) error {
	m.tags[tag.Id] = &tag
	return nil
}

func (m Repository) GetTagById(id string) (*domain2.Tag, error) {
	tag, ok := m.tags[id]
	if !ok {
		return nil, errors.New("tag not found")
	}
	return tag, nil
}

func (m Repository) GetAllTags() ([]*domain2.Tag, error) {
	tags := make([]*domain2.Tag, len(m.tags))
	for _, tag := range m.tags {
		tags = append(tags, tag)
	}
	return tags, nil
}

func (m Repository) GetTagByTitle(title string) (*domain2.Tag, error) {
	for _, tag := range m.tags {
		if tag.Title == title {
			return tag, nil
		}
	}
	return nil, errors.New("tag not found")
}

func (m Repository) GetTagsForSound(soundId string) ([]*domain2.Tag, error) {
	tagIds := make([]string, 0)
	for _, item := range m.soundTags {
		if item.soundId == soundId {
			tagIds = append(tagIds, item.tagId)
		}
	}
	tags := make([]*domain2.Tag, 0)

	for _, id := range tagIds {
		tag, ok := m.tags[id]
		if !ok {
			return nil, errors.New(fmt.Sprintf("tag with id %s doesn't exist", id))
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (m Repository) GetSoundById(id string) (*domain2.Sound, error) {
	sound, ok := m.sounds[id]
	if !ok {
		return nil, errors.New("sound not found")
	}

	return sound, nil
}

func (m Repository) GetAllSounds() ([]*domain2.Sound, error) {
	sounds := make([]*domain2.Sound, 0)
	for _, value := range m.sounds {
		sounds = append(sounds, value)
	}
	return sounds, nil
}

func (m Repository) GetSounds(limit, offset int) ([]*domain2.Sound, error) {
	sounds := make([]*domain2.Sound, 0)

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

func (m Repository) GetSoundsNameLike(name string) ([]*domain2.Sound, error) {
	sounds := make([]*domain2.Sound, 0)
	for _, value := range m.sounds {
		if strings.Contains(value.Name, name) {
			sounds = append(sounds, value)
		}
	}
	return sounds, nil
}

func (m Repository) GetSoundsByTagId(tagId string) ([]*domain2.Sound, error) {
	sounds := make([]*domain2.Sound, 0)

	for _, sound := range m.sounds {
		for _, tag := range sound.Tags {
			if tag.Id == tagId {
				sounds = append(sounds, sound)
			}
		}
	}

	return sounds, nil
}

func (m Repository) GetSoundsByVehicleId(vehicleId string) ([]*domain2.Sound, error) {
	sounds := make([]*domain2.Sound, 0)

	for _, sound := range m.sounds {
		if sound.VehicleId == vehicleId {
			sounds = append(sounds, sound)
		}
	}

	return sounds, nil
}
