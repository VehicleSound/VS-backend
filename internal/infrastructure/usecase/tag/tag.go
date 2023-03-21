package tag

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/utils"
	"github.com/timickb/transport-sound/internal/interfaces"
)

type UseCase struct {
	repo usecase.Repository
	log  interfaces.Logger
}

func NewTagUseCase(repo usecase.Repository, log interfaces.Logger) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreateTag(name string) (string, error) {
	if !utils.ValidateTag(name) {
		return "", errors.New("err create tag title too short")
	}

	existing, err := u.repo.GetTagByTitle(name)
	if err == nil && existing.Title == name {
		return "", errors.New("err create tag: title already exists")
	}

	tagId := uuid.NewString()
	tag := domain.Tag{
		Id:    tagId,
		Title: name,
	}

	if err := u.repo.CreateTag(tag); err != nil {
		return "", fmt.Errorf("err create tag: %w", err)
	}

	return tagId, nil
}

func (u *UseCase) GetTagById(id string) (*domain.Tag, error) {
	tag, err := u.repo.GetTagById(id)
	if err != nil {
		return nil, fmt.Errorf("err get tag by id: %w", err)
	}

	return tag, nil
}

func (u *UseCase) GetTagByTitle(title string) (*domain.Tag, error) {
	tag, err := u.repo.GetTagByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("err get tag by title: %w", err)
	}

	return tag, nil
}

func (u *UseCase) GetAllTags() ([]*domain.Tag, error) {
	tags, err := u.repo.GetAllTags()
	if err != nil {
		return nil, fmt.Errorf("err get all tags: %w", err)
	}

	return tags, nil
}
